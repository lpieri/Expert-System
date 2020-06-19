package main

import (
	"os"
)

func main() {
	var file sFile
	file = checkArgs(os.Args)
	resolve(file)
	printResult(file)
}
