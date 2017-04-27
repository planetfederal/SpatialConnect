package main

//Schema
type Schema interface {
	build()
	test()
	deploy()
}
