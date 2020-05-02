package cmd

import (
	"log"
	"os/exec"
)

func RunCmd(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
