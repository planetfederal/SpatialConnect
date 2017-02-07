package main

// Repo
type Repo interface {
	checkout()
	clean()
	rebase(remote string, branch string)
}
