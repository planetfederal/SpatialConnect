package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func combineArgsRecur(acc []string, c []string) []string {
	if len(c) == 0 {
		return acc
	} else if strings.Contains(c[0], "'") {
		return append(acc, c[0])
	} else {
		return combineArgsRecur(append(acc, c[0]), c[1:])
	}
}

func combineQuotesRecur(acc []string, c []string) []string {
	if len(c) == 0 {
		return acc
	} else if strings.Contains(c[0], "'") {
		newArr := append([]string{}, c[0])
		combinedQuote := combineArgsRecur(newArr, c[1:])
		combinedArgs := append(acc, strings.Join(combinedQuote, " "))
		return combineQuotesRecur(combinedArgs, c[0+len(combinedQuote):])
	} else {
		return combineQuotesRecur(append(acc, c[0]), c[1:])
	}
}

func combineQuotes(c []string) []string {
	return combineQuotesRecur([]string{}, c)
}

func executeCmd(c string) *exec.Cmd {
	args := strings.Split(c, " ")
	newArgs := combineQuotes(args[1:])
	fmt.Printf("%s %s\n", args[0], strings.Join(newArgs, "_"))
	command := exec.Command(args[0], newArgs...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command
}

func executeCmds(cmds []string) {
	for i := 0; i < len(cmds); i++ {
		executeCmd(cmds[i])
	}
}
