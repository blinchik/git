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

}
