package gl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"sync"

	"github.com/xanzy/go-gitlab"
	"golang.org/x/crypto/ssh"
	gitgo "gopkg.in/src-d/go-git.v4"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	gitgoSSH "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// getSshKeyAuth read private key and return AuthMethod
func getSshKeyAuth(privateSshKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, err := ioutil.ReadFile(privateSshKeyFile)
	if err != nil {
		log.Fatal(err)

	}

	signer, err := ssh.ParsePrivateKey([]byte(sshKey))
	if err != nil {
		log.Fatal(err)

	}

	auth = &gitgoSSH.PublicKeys{User: "git", Signer: signer}
	return auth
}

// GroupCloneAllProjects clones all gitlab projects in the given group and it's subgroups
func GroupCloneAllProjects(groupID, key string) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token, empty := os.LookupEnv("GITLAB_PRIVATE_TOKEN")

	if !empty {
		log.Fatal("There is no private token. Please set 'GITLAB_PRIVATE_TOKEN' environment variable")
	}

	git, err := gitlab.NewClient(token)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var cloneProjects func(groupID, fullPath string, wg *sync.WaitGroup)
	var searchSubgroups func(groupID string)

	// clone all projects in a given group
	cloneProjects = func(groupID, fullPath string, wg *sync.WaitGroup) {

		defer wg.Done()

		opt := &gitlab.ListGroupProjectsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
				Page:    1,
			}}

		projects, _, err := git.Groups.ListGroupProjects(groupID, opt)

		if err != nil {
			log.Fatal(err)

		}

		usr, err := user.Current()

		if err != nil {
			log.Fatal(err)

		}

		for _, i := range projects {
			fmt.Printf("Cloning %s into %s \n", i.Path, fullPath)

			repoLocalPath := fmt.Sprintf("%s/%s\n", fullPath, i.Path)

			_, err := gitgo.PlainClone(repoLocalPath, false, &gitgo.CloneOptions{

				Auth:     getSshKeyAuth(usr.HomeDir + key),
				URL:      i.SSHURLToRepo,
				Progress: os.Stdout,
			})

			if err != nil {
				log.Fatal(err)

			}

		}

	}

	// get all subgroups within a given group and create matching local directories.
	searchSubgroups = func(groupID string) {

		var wg sync.WaitGroup

		opt := &gitlab.ListSubgroupsOptions{}
		groups, _, err := git.Groups.ListSubgroups(groupID, opt)

		if err != nil {
			log.Fatal(err)

		}

		for _, i := range groups {
			os.MkdirAll(i.FullPath, os.ModePerm)

			searchSubgroups(strconv.Itoa(i.ID))

			wg.Add(1)
			go cloneProjects(strconv.Itoa(i.ID), i.FullPath, &wg)
		}

		wg.Wait()

	}

	searchSubgroups(groupID)

}
