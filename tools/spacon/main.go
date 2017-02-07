package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	proj := os.Args[1]
	command := os.Args[2]

	cfg := readConfigFile()

	ios := IOSSDK{org: cfg.Origin, name: cfg.Upstream}
	p := make(map[string]IOSSDK)

	p["ios"] = ios
	//
	var r Repo
	r = ios

	switch proj {
	case "init":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What is your origin remote alias?")
		origin, _ := reader.ReadString('\n')
		fmt.Println("What is your upstream remote alias?")
		upstream, _ := reader.ReadString('\n')
		c := Configuration{strings.TrimRight(origin, "\n"), strings.TrimRight(upstream, "\n")}
		writeConfigFile(c)
	case "all":
		fmt.Print("All")
	case "ios":
		fmt.Println("ios")
		r.checkout()
	case "android":
		fmt.Println("android")
	case "js":
		fmt.Println("js")
	case "server":
		fmt.Println("server")
	case "mobile":
		fmt.Println("mobile")
	default:
		fmt.Println("Unknown command: %s", proj)
	}

}
