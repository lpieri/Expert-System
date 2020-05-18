package main

import "fmt"

func printError(err error) {
	fmt.Println("Error: expert-system")
	fmt.Println(err)
}

func printUsage() {
	fmt.Println("Usage: expert-system [options]")
	fmt.Println(" --input -i [file]\tInput file")
}
