package main

import (
	"fmt"
	"strings"
)

func serchForParentheses(rule string) (int, int) {
	println("rule in parentheses = ", rule)
	first := -1
	last := -1
	for i := 0; i < len(rule); i++ {
		println("i = ", i, "  rule[i] = ", string(rule[i]))
		if rule[i] == '(' && first == -1 {
			first = i
		} else if rule[i] == ')' && i > first {
			if first == -1 {
				printErrorMsg("Error in input file, please check the parentheses")
			}
			last = i
		}
	}
	println("first = ", first, "   last = ", last)
	if first != 0 && last != len(rule)-1 {
		if !(strings.Contains("+|^", string(rule[first-1])) || (strings.Contains("+|^", string(rule[last+1])))) {
			printErrorMsg("Error in input file, please check the parentheses2")
		}
	}
	return first, last
}

func checkForSymbol(t *Tree, rule string, symb string) *Tree {
	println("rule = ", rule, "symb =", symb)
	tab := strings.Split(rule, symb)
	if t == nil {
		t = &Tree{nil, symb, nil}
	}
	t.Value = symb
	t.Left = isPrio(t.Left, strings.TrimSpace(tab[0]))
	t.Right = isPrio(t.Right, strings.TrimSpace(tab[1]))
	println("t.Value =", t.Value)
	println("t.Left =", treeToString(t.Left))
	println("t.Right =", treeToString(t.Right))
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
		println("first, last  ==", first, ",", last) // faire des trucs avec les parenthese
		// prendre rule[first-1:last]A | (B + C) + D
		println("last = ", last, "len de rule = ", len(rule))
		if first == 0 && last == len(rule)-1 {
			println("cas1")
			rule = rule[1 : len(rule)-1]
			t = isPrio(t, rule)
		} else if first == 0 {
			println("cas2")
			//t = checkForSymbol(t, rule, string(rule[last+1]))
			if t == nil {
				t = &Tree{nil, string(rule[last+1]), nil}
			}
			t.Value = string(rule[last+1])
			t.Left = isPrio(t.Left, strings.TrimSpace(rule[:last+1]))
			t.Right = isPrio(t.Right, strings.TrimSpace(rule[last+2:]))
		} else if last == len(rule)-1 {
			println("cas3")
			//t = checkForSymbol(t, rule, string(rule[first-1]))
			if t == nil {
				t = &Tree{nil, string(rule[first-1]), nil}
			}
			t.Value = string(rule[first-1])
			t.Left = isPrio(t.Left, strings.TrimSpace(rule[:first-1]))
			t.Right = isPrio(t.Right, strings.TrimSpace(rule[first:]))
		} else if mCompDict[string(rule[first-1])] < mCompDict[string(rule[last+1])] { //si operateur AVANT parenthe est moins prioritaire
			println("cas4")
			//t = checkForSymbol(t, rule, string(rule[first-1]))
			if t == nil {
				t = &Tree{nil, string(rule[first-1]), nil}
			}
			t.Value = string(rule[first-1])
			t.Left = isPrio(t.Left, strings.TrimSpace(rule[:first-1]))
			t.Right = isPrio(t.Right, strings.TrimSpace(rule[first:]))
		} else {
			println("cas5")
			//t = checkForSymbol(t, rule, string(rule[last+1]))
			if t == nil {
				t = &Tree{nil, string(rule[last+1]), nil}
			}
			t.Value = string(rule[last+1])
			t.Left = isPrio(t.Left, strings.TrimSpace(rule[:last+1]))
			t.Right = isPrio(t.Right, strings.TrimSpace(rule[last+2:]))
		}
	} else if len(strings.TrimSpace(rule)) == 1 {
		if t == nil {
			t = &Tree{nil, strings.TrimSpace(rule), nil}
		}
		println("rule in fucking isPRio", rule)
		t.Value = strings.TrimSpace(rule)
	} else if len(strings.TrimSpace(rule)) == 2 {
		if !strings.Contains(rule, "!") {
			printErrorMsg("eror syntax please check the input file")
		}
		if t == nil {
			t = &Tree{nil, strings.TrimSpace(rule), nil}
		}
		println("rule in fucking isPRio", rule)
		t.Value = strings.TrimSpace(rule)
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
