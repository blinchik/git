package gl

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/xanzy/go-gitlab"
)

// GetSubgroups imports subgroups under specific group
func GetSubgroups(token, groupID string) ([]*gitlab.Group, []string) {

	var subgroupsID []string

	git, err := gitlab.NewClient(token)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	opt := &gitlab.ListSubgroupsOptions{}
	groups, resp, err := git.Groups.ListSubgroups(groupID, opt)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	log.Print(string(bodyBytes))

	for _, i := range groups {
		subgroupsID = append(subgroupsID, strconv.Itoa(i.ID))
		os.MkdirAll(i.FullPath, os.ModePerm)
	}
	return groups, subgroupsID

}
