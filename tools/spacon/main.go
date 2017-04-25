package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleAllCommand(proj map[string]Repo, cmd string) {
	switch cmd {
	case "clone":
		for k, v := range proj {
			fmt.Printf("Cloning %s\n", k)
			var p Project
			p = v
			p.clone()
		}
	case "clean":
		for _, v := range proj {
			var p Project
			p = v
			p.clean()
		}
	case "checkout":
		for _, v := range proj {
			var p Project
			p = v
			p.checkout(os.Args[3])
		}
	default:
		panic("Command:" + cmd + " ")
	}

}

func handleCommand(p Project, cmd string) {
	switch cmd {
	case "build":
		p.build()
	case "run":
		p.run()
	case "clean":
		p.clean()
	case "clone":
		p.clone()
	case "checkout":
		p.checkout(os.Args[3])
	case "test":
		p.test()
	case "deploy":
		p.deploy()
	default:
		fmt.Print("command not found\n")

	}
}

func main() {
	cfg := readConfigFile()

	if os.Args[1] == "init" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What is your origin remote alias?")
		origin, _ := reader.ReadString('\n')
		fmt.Println("What is your upstream remote alias?")
		upstream, _ := reader.ReadString('\n')
		c := Configuration{strings.TrimRight(origin, "\n"), strings.TrimRight(upstream, "\n")}
		writeConfigFile(c)
		return
	}

	ios := Repo{cfg, "spatialconnect-ios-sdk"}
	android := Repo{cfg, "spatialconnect-android-sdk"}
	js := Repo{cfg, "spatialconnect-js"}
	server := Repo{cfg, "spatialconnect-server"}
	schema := Repo{cfg, "spatialconnect-schema"}
	mobile := Repo{cfg, "spatialconnect-mobile"}

	p := make(map[string]Repo)

	p["ios"] = ios
	p["android"] = android
	p["js"] = js
	p["server"] = server
	p["schema"] = schema
	p["mobile"] = mobile

	ca2 := os.Args[1]
	proj, ok := p[ca2]
	command := os.Args[2]

	if ca2 == "all" {
		handleAllCommand(p, command)
		return
	} else if !ok {
		panic("You need to select a project [ios|android|js|server|schema|mobile|all]")
	}

	if len(os.Args) < 3 {
		panic("You don't have a command. $>spacon <project> <command>")
	}
	var sp Project
	sp = proj
	handleCommand(sp, command)
}
