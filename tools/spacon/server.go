package main

//Server
type Server interface {
	build()
	run()
	test()
	release()
}
