package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

func isPrio(rule []string) bool {
	
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
