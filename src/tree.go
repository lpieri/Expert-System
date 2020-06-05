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

func serchForParentheses(rule string) (int, int) {
	println("rule in parentheses = ", rule)
	first := -1
	last := -1
	for i := 0; i < len(rule); i++ {
		println("i = ", i, "  rrule[i] = ", string(rule[i]))
		if rule[i] == '(' {
			first = i
		} else if rule[i] == ')' && i > first {
			if first == -1 {
				printErrorMsg("Error in input file, please check the parentheses")
			}
			last = i
		}
	}
	if !(strings.Contains("+|^", string(rule[first-1]))) {
		printErrorMsg("Error in input file, please check the parentheses2")
	}
	return first, last
}

func isPrio(t *Tree, rule string) *Tree {
	if strings.Contains(rule, "^") && !strings.Contains(rule, "(") {
		tab := strings.Split(rule, "^")
		println("tab[0] =", strings.TrimSpace(tab[0]))
		println("tab[1] =", strings.TrimSpace(tab[1]))
		t.Value = "^"
		t.Left = isPrio(t.Left, strings.TrimSpace(tab[0]))
		t.Right = isPrio(t.Right, strings.TrimSpace(tab[1]))
	} else if strings.Contains(rule, "|") && !strings.Contains(rule, "(") {
		tab := strings.Split(rule, "|")
		println("tab[0] =", strings.TrimSpace(tab[0]))
		println("tab[1] =", strings.TrimSpace(tab[1]))
		t.Value = "|"
		t.Left = isPrio(t.Left, strings.TrimSpace(tab[0]))
		t.Right = isPrio(t.Right, strings.TrimSpace(tab[1]))
	} else if strings.Contains(rule, "+") && !strings.Contains(rule, "(") {
		tab := strings.Split(rule, "+")
		println("tab[0] =", strings.TrimSpace(tab[0]))
		println("tab[1] =", strings.TrimSpace(tab[1]))
		t.Value = "+"
		t.Left = isPrio(t.Left, strings.TrimSpace(tab[0]))
		t.Right = isPrio(t.Right, strings.TrimSpace(tab[1]))
	} else if strings.Contains(rule, "()") {
		i, j := serchForParentheses(rule)
		print("i, j  == ", i, " , ", j) // faire des trucs avec les parenthese
	} else {
		println("rule in fucking isPRio", rule)
		//t.Value = strings.TrimSpace(rule)
	}
	return t
}

func newTree(rule []string) *Tree {
	var t *Tree
	// for i := 0; i < len(rule); i++ {
	// 	elem := rule[i]
	// 	print(elem)
	// 	t = insert(t, elem)
	// }

	concat := strings.Join(rule, "")
	fmt.Println("concat = ", concat)
	depart := isPrio(t, concat)
	println("depart = ", treeToString(depart))
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
