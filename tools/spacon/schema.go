package main

//Schema
type Schema interface {
	build()
	run()
	test()
	release()
}
