package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	curPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	chgDirCmd := exec.Command("cmd", "cd "+curPath+"../")

	if err := chgDirCmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	newPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	runExec := exec.Command("cmd", "start StartGoPractice "+newPath+" /min practice.exe")

	if err := runExec.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
