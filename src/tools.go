package main

import (
	"fmt"
	"os"
)

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
