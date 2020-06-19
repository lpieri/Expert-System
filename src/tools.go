package main

import (
	"fmt"
	"os"
)

func removeIndexFormTab(i int, tab []string) []string {
	copy(tab[i:], tab[i+1:])
	tab[len(tab)-1] = ""
	tab = tab[:len(tab)-1]
	return tab
}

func printErrorMsg(msg string) {
	fmt.Println("Error: expert-system")
	fmt.Println(msg)
	os.Exit(1)
}

func printError(err error) {
	fmt.Println("Error: expert-system")
	fmt.Println(err)
	os.Exit(1)
}

func printUsage() {
	fmt.Println("Usage: expert-system [options]")
	fmt.Println(" --input -i [file]\tInput file")
	os.Exit(0)
}

func printResult(file sFile) {
	for i := 0; i < len(file.Queries); i++ {
		if vars[string(file.Queries[i])] == "" {
			vars[string(file.Queries[i])] = "false"
		}
		fmt.Println(file.Queries[i], "is", vars[string(file.Queries[i])])
	}
}

func delChar(s string, index int) string {
	return s[0:index] + s[index+1:]
}
