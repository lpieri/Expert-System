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
