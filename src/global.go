package main

/*Variables Gloabales*/
var vars = map[string]string{}

type sRule struct {
	Facts      []string
	Conclusion []string
}

type sFile struct {
	Rules   []sRule
	Queries []string
	Init    []string
}

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

var mCompDict = map[string]int{"^": 0, "|": 1, "+": 2, "": 10}
