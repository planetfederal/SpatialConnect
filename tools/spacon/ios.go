package main

import (
	"os"
	"os/exec"
	"syscall"
)

func path(i IOSSDK) string {
	return i.org + "/" + i.name
}

func (i IOSSDK) checkout(branch string) {
	err := executeCmds([]string{"cd " + i.name, "git checkout -B " + branch})
	if err != nil {
		panic(err)
	}
}

func (i IOSSDK) clone() {
	execErr := executeCmd("git clone git@github.com:" + path(i) + ".git")
	if execErr != nil {
		panic(execErr)
	}
}

func (i IOSSDK) clean() {
	args := []string{"rm", i.name}
	env := os.Environ()
	binary, _ := exec.LookPath("rm")
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func (i IOSSDK) rebase(remote, branch string) {
	args := []string{"git", "pull", "--rebase", remote + "/" + branch}
	env := os.Environ()
	binary, _ := exec.LookPath("git")
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
