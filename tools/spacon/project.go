package main

import (
	"fmt"
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

	install()
	build()
	run()
	unitTest()
	integrationTest()
	release()
}

func path(r Repo) string {
	return r.cfg.Origin + "/" + r.name
}

func (r Repo) checkout(branch string) {
	fmt.Printf("Checking out %s in %s\n", branch, r.name)
	os.Chdir("./" + r.name)
	err := executeCmds([]string{"git fetch --all", "git checkout -B " + branch + " origin/" + branch})
	if err != nil {
		fmt.Printf("Error checking out branch:%s\n", branch)
	}
	os.Chdir("..")
}

func (r Repo) clone() {
	err := executeCmd("git clone git@github.com:" + path(r) + ".git")
	if err != nil {
		fmt.Printf("Error cloning %s\n", path(r))
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = executeCmd("git --git-dir=./" + r.name + "/.git remote add upstream git@github.com:" +
		r.cfg.Upstream + "/" + r.name + ".git")
	if err != nil {
		fmt.Printf("Can't add remote alias:%s/%s\n", r.cfg.Upstream, r.name)
	}
}

func (r Repo) clean() {
	err := executeCmd("rm -rf " + r.name)
	if err != nil {
		fmt.Printf("Can't remove the repo:%s\n", r.name)
	}
}

func (r Repo) install() {
	pc := readProjectConfigFile(r)
	os.Chdir("./" + r.name)
	err := executeCmd(pc.Install)
	if err != nil {
		fmt.Printf("Can't install")
	}
	os.Chdir("./..")
}

func (r Repo) build() {
	pc := readProjectConfigFile(r)
	os.Chdir("./" + r.name)
	err := executeCmd(pc.Build)
	if err != nil {
		fmt.Printf("Can't build")
	}
	os.Chdir("./..")
}

func (r Repo) run() {
	pc := readProjectConfigFile(r)
	err := executeCmd(pc.Run)
	if err != nil {
		fmt.Printf("Can't Run")
	}
	os.Chdir("./..")
}

func (r Repo) unitTest() {
	pc := readProjectConfigFile(r)
	os.Chdir("./" + r.name)
	err := executeCmd(pc.UnitTest)
	if err != nil {
		fmt.Printf("Can't Unit Test")
	}
	os.Chdir("./..")
}

func (r Repo) integrationTest() {

}

func (r Repo) release() {

}
