package main

import (
	"fmt"
	"io/ioutil"
)

func openFile(fileName string) {
	bits, err := ioutil.ReadFile(fileName)
	if err != nil {
		printError(err)
	} else {
		data := string(bits)
		fmt.Println(data)
	}
}

func checkArgs(arguments []string) {
	if len(arguments) < 3 {
		printUsage()
	} else {
		for i := 1; i < len(arguments); i++ {
			if arguments[i] == "--input" || arguments[i] == "-i" {
				openFile(arguments[i+1])
			} else {
				printUsage()
			}
		}
	}
}
