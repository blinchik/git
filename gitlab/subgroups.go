package gl

//
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//	"strconv"
//
//	"github.com/xanzy/go-gitlab"
//)
//
//// GetSubgroups imports subgroups under specific group
//func GetSubgroups(token, groupID string) ([]*gitlab.Group, []string) {
//
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//
//	var subgroupsID []string
//
//	git, err := gitlab.NewClient(token)
//
//	if err != nil {
//		log.Fatalf("Failed to create client: %v", err)
//	}
//
//	subgroupsID := subgroups(git, groupID)
//
//	return groups, subgroupsID
//
//}
//
//func subgroups(git *gitlab.Client, groupID string) []string {
//
//	var subgroupsID []string
//
//	opt := &gitlab.ListSubgroupsOptions{}
//	groups, resp, err := git.Groups.ListSubgroups(groupID, opt)
//
//	if err != nil {
//		log.Fatal(err)
//
//	}
//
//	defer resp.Body.Close()
//	bodyBytes, err := ioutil.ReadAll(resp.Body)
//	log.Print(string(bodyBytes))
//
//	for _, i := range groups {
//		subgroupsID = append(subgroupsID, strconv.Itoa(i.ID))
//		os.MkdirAll(i.FullPath, os.ModePerm)
//		fmt.Println(i.FullPath, i.Name)
//	}
//
//return 	subgroupsID
//
//}
//
////func gitlabProjects(git *gitlab.client){
////
////
////	opt := &gitlab.ListGroupProjectsOptions{}
////	projects, resp, err := git.Groups.ListGroupProjects(groupID, opt)
////
////	if err != nil {
////		log.Fatal(err)
////
////	}
////
////	for _, i := range projects {
////		subgroupsID = append(subgroupsID, strconv.Itoa(i.ID))
////		os.MkdirAll(i.FullPath, os.ModePerm)
////	}
////
////
////
////	defer resp.Body.Close()
////	bodyBytes, err := ioutil.ReadAll(resp.Body)
////	log.Print(string(bodyBytes))
////
////
////}
