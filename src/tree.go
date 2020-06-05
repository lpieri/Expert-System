package main

import (
	"fmt"
	"strings"
)

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

func serchForParentheses(rule []string) (int, int) {
	println("rule in parentheses = ", rule)
	first := -1
	last := -1
	for i := 0; i < len(rule); i++ {
		println("i = ", i, "  rrule[i] = ", rule[i])
		if strings.Contains(rule[i], "(") {
			first = i
		} else if strings.Contains(rule[i], ")") && i > first && first != -1 {
			last = i
		} else {
			printErrorMsg("Error in input file, please check the parentheses")
		}
	}
	return first, last
}

func isPrio(rule []string) bool {
	i, j := serchForParentheses(rule)
	println("i, j = ", i, " , ", j)
	for i := i; i <= j; i++ {
		
	}
	return true
}

// func insert(t *Tree, v string) *Tree {
// 	if t == nil {
// 		return &Tree{nil, v, nil}
// 	}
// 	if isPrio(v, t.Value) {
// 		t.Left = insert(t.Left, v)
// 	} else {
// 		t.Right = insert(t.Right, v)
// 	}
// 	return t
// }

func newTree(rule []string) *Tree {
	var t *Tree
	// for i := 0; i < len(rule); i++ {
	// 	elem := rule[i]
	// 	print(elem)
	// 	t = insert(t, elem)
	// }
	depart := isPrio(rule)
	println(depart)
	return t
}

func treeToString(t *Tree) string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += treeToString(t.Left) + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + treeToString(t.Right)
	}
	return "(" + s + ")"
}
