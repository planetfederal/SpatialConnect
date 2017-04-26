package main

import (
	"fmt"
	"go/scanner"
	"os"
)

//Repo for holding data
type Repo struct {
	cfg  Configuration
	name string
}

//Project for managing repo
type Project interface {
	checkout(branch string)
	clean()
	clone()
	build()
	run()
	test()
	deploy()
}

func path(r Repo) string {
	return r.cfg.Origin + "/" + r.name
}

func (r Repo) checkout(branch string) {
	fmt.Printf("Checking out %s in %s\n", branch, r.name)
	os.Chdir("./" + r.name)
	executeCmds([]string{"git fetch --all", "git checkout -B " + branch + " origin/" + branch})
	os.Chdir("..")
}

func (r Repo) clone() {
	err := executeCmd("git clone git@github.com:" + path(r) + ".git").Run()
	if err != nil {
		scanner.PrintError(os.Stderr, err)
		fmt.Printf("Error cloning %s\n", path(r))
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = executeCmd("git --git-dir=./" + r.name + "/.git remote add upstream git@github.com:" +
		r.cfg.Upstream + "/" + r.name + ".git").Run()
	if err != nil {
		scanner.PrintError(os.Stderr, err) // <-- here
		fmt.Printf("Can't add remote alias:%s/%s\n", r.cfg.Upstream, r.name)
	}
}

func (r Repo) clean() {
	err := executeCmd("rm -rf " + r.name).Run()
	if err != nil {
		scanner.PrintError(os.Stderr, err) // <-- here
		fmt.Printf("Can't remove the repo:%s\n", r.name)
	}
}

func (r Repo) build() {
	pc := readProjectConfigFile(r)
	os.Chdir("./" + r.name)
	err := executeCmd(pc.Build).Run()
	if err != nil {
		scanner.PrintError(os.Stderr, err)
	}
	os.Chdir("./..")
}

func (r Repo) run() {
	pc := readProjectConfigFile(r)
	err := executeCmd(pc.Run).Run()
	if err != nil {
		scanner.PrintError(os.Stderr, err)
	}
	os.Chdir("./..")
}

func (r Repo) test() {
	pc := readProjectConfigFile(r)
	os.Chdir("./" + r.name)
	err := executeCmd(pc.Test).Run()
	if err != nil {
		scanner.PrintError(os.Stderr, err)
	}
	os.Chdir("./..")
}

func (r Repo) deploy() {
	func (r Repo) deploy() {
	    pc := readProjectConfigFile(r)
	    os.Chdir("./" + r.name)
	    err := executeCmd(pc.Deploy)
	    if err != nil {
	        scanner.PrintError(os.Stderr, err)
	    }
	    os.Chdir("./..")
	}
}
