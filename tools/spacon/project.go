package main

// IOSSDK
type program struct {
	org           string
	name          string
	defaultBranch string
}

type Project interface {
	build()
	run()
	test()
	release()
}
