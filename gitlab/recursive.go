package gl

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/xanzy/go-gitlab"
	gitgo "gopkg.in/src-d/go-git.v4"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// GroupCloneAllProjects clones all gitlab projects in the given group and it's subgroups
func GroupCloneAllProjects(token, groupID string) {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	git, err := gitlab.NewClient(token)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var cloneProjects func(groupID, fullPath string)
	var searchSubgroups func(groupID string)

	cloneProjects = func(groupID, fullPath string) {

		opt := &gitlab.ListGroupProjectsOptions{}
		projects, _, err := git.Groups.ListGroupProjects(groupID, opt)

		if err != nil {
			log.Fatal(err)

		}

		for _, i := range projects {
			fmt.Printf("Cloning %s into %s \n", i.Path, fullPath)

			repoLocalPath := fmt.Sprintf("%s/%s", fullPath, i.Path)

			_, err := gitgo.PlainClone(repoLocalPath, false, &gitgo.CloneOptions{

				Auth: &http.BasicAuth{
					Username: "go bot",
					Password: token,
				},
				URL:      i.HTTPURLToRepo,
				Progress: os.Stdout,
			})

			if err != nil {
				log.Fatal(err)

			}

		}

	}

	searchSubgroups = func(groupID string) {

		opt := &gitlab.ListSubgroupsOptions{}
		groups, _, err := git.Groups.ListSubgroups(groupID, opt)

		if err != nil {
			log.Fatal(err)

		}

		for _, i := range groups {
			os.MkdirAll(i.FullPath, os.ModePerm)

			searchSubgroups(strconv.Itoa(i.ID))
			cloneProjects(strconv.Itoa(i.ID), i.FullPath)
		}

	}

	searchSubgroups(groupID)

}
