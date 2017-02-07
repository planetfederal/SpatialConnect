package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func executeCmd(cmd string) error {
	args := strings.Split(cmd, " ")

	env := os.Environ()
	binary, _ := exec.LookPath(args[0])
	execErr := syscall.Exec(binary, args, env)
	return execErr
}

func executeCmds(cmd []string) error {
	var err error
	for i := 0; i < len(cmd); i++ {
		err = executeCmd(cmd[i])
	}
	return err
}
