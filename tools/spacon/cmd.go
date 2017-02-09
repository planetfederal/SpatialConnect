package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeCmd(cmd string) error {
	args := strings.Split(cmd, " ")
	fmt.Printf("%s\n", args)
	return exec.Command(args[0], args[1:]...).Run()
}

func executeCmds(cmd []string) error {
	var err error
	for i := 0; i < len(cmd); i++ {
		err = executeCmd(cmd[i])
	}
	return err
}
