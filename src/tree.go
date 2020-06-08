package main

import (
	"fmt"
	"strings"
)

func serchForParentheses(rule string) (int, int) {
	first := -1
	last := -1
	for i := 0; i < len(rule); i++ {
		if rule[i] == '(' && first == -1 {
			first = i
		} else if rule[i] == ')' && i > first {
			if first == -1 {
				printErrorMsg("Error in input file, please check the parentheses")
			}
			last = i
		}
	}
	if first != 0 && last != len(rule)-1 {
		if !(strings.Contains("+|^", string(rule[first-1])) || (strings.Contains("+|^", string(rule[last+1])))) {
			printErrorMsg("Error in input file, please check the parentheses2")
		}
	}
	return first, last
}

func checkForSymbol(t *Tree, rule string, symb string) *Tree {
	tab := strings.Split(rule, symb)
	if t == nil {
		t = &Tree{nil, symb, nil}
	}
	t.Value = symb
	t.Left = isPrio(t.Left, strings.TrimSpace(tab[0]))
	t.Right = isPrio(t.Right, strings.TrimSpace(tab[1]))
	return t
}

func sliceParenthese(t *Tree, rule string, j int, i int) *Tree {
	if t == nil {
		t = &Tree{nil, string(rule[j]), nil}
	}
	t.Value = string(rule[j])
	t.Left = isPrio(t.Left, strings.TrimSpace(rule[:j]))
	t.Right = isPrio(t.Right, strings.TrimSpace(rule[i:]))
	return t
}

func isPrio(t *Tree, rule string) *Tree {
	if strings.Contains(rule, "^") && !strings.ContainsAny(rule, "()") {
		t = checkForSymbol(t, rule, "^")
	} else if strings.Contains(rule, "|") && !strings.ContainsAny(rule, "()") {
		t = checkForSymbol(t, rule, "|")
	} else if strings.Contains(rule, "+") && !strings.ContainsAny(rule, "()") {
		t = checkForSymbol(t, rule, "+")
	} else if strings.ContainsAny(rule, "()") {
		first, last := serchForParentheses(rule)
		if first == 0 && last == len(rule)-1 {
			rule = rule[1 : len(rule)-1]
			t = isPrio(t, rule)
		} else if first == 0 {
			t = sliceParenthese(t, rule, last+1, last+2)
		} else if last == len(rule)-1 {
			t = sliceParenthese(t, rule, first-1, first)
		} else if mCompDict[string(rule[first-1])] < mCompDict[string(rule[last+1])] {
			t = sliceParenthese(t, rule, first-1, first)
		} else {
			t = sliceParenthese(t, rule, last+1, last+2)
		}
	} else if len(strings.TrimSpace(rule)) == 1 {
		if t == nil {
			t = &Tree{nil, strings.TrimSpace(rule), nil}
		}
		t.Value = strings.TrimSpace(rule)
	} else if len(strings.TrimSpace(rule)) == 2 {
		if !strings.Contains(rule, "!") {
			printErrorMsg("eror syntax please check the input file")
		}
		if t == nil {
			t = &Tree{nil, strings.TrimSpace(rule), nil}
		}
		t.Value = strings.TrimSpace(rule)
	}
	return t
}

func newTree(rule []string) *Tree {
	var t *Tree
	concat := strings.Join(rule, "")
	depart := isPrio(t, concat)
	return depart
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
