package main

import (
	gl "cloneall/gitlab"
	"log"
	"os"
)

func main() {
	token, empty := os.LookupEnv("GITLAB_PRIVATE_TOKEN")

	if !empty {
		log.Fatal("There is no private token. Please set 'GITLAB_PRIVATE_TOKEN' environment variable")
	}

	gl.GroupCloneAllProjects(token, "10299617")
	//	x, _ := json.MarshalIndent(subgroups, "", "    ")
	//	fmt.Println(string(x), subgroupsID)

	//	rec(subgroupsID, token)

}

//func rec(subgroupsID []string, token string) {
//
//	if len(subgroupsID) != 0 {
//		for _, i := range subgroupsID {
//
//			fmt.Println(i)
//
//			g, subgroupsID := gl.GetSubgroups(token, i)
//			x, _ := json.MarshalIndent(g, "", "    ")
//			fmt.Println(string(x), subgroupsID)
//
//			rec(subgroupsID, token)
//		}
//	}
//
//}
