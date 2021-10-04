package main

import (
	gl "blinchikgit/gitlab"
	"flag"
	"os"
)

func main() {

	glclone := flag.Bool("gl", false, "gl clone Gitlab projects under a specific group")
	flag.Parse()

	if *glclone {
		gl.GroupCloneAllProjects(os.Args[2], os.Args[3])
	}

}
