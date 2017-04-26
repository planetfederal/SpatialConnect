package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func executeCmd(c string) *exec.Cmd {
	args := strings.Split(c, " ")
	argStr := strings.Join(args[1:], " ")
	fmt.Printf("%s\n", args)
	command := exec.Command(args[0], argStr)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command
}

func executeCmds(cmds []string) {
	for i := 0; i < len(cmds); i++ {
		executeCmd(cmds[i])
	}
}
