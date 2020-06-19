package main

import (
	"strconv"
	"strings"
)

var gRules []sRule

func fillTree(toTree []string) *Tree {
	tree := newTree(toTree)
	return tree
}

func backtraking(newLetter string) string {
	tab := checkLetterInConc(newLetter, gRules)
	for i := 0; i < len(tab); i++ {
		newFact := fillTree(gRules[tab[i]].Facts)
		newValue := browseTree(newFact)
		if newValue == "" {
			return "false"
		}
		if newValue == "true" {
			vars[newLetter] = newValue
		}
		return newValue
	}
	return "false"
}

func operatorSolver(leftVal string, ope string, rightVal string) string {
	res := ""
	left, err := strconv.ParseBool(leftVal)
	right, err2 := strconv.ParseBool(rightVal)
	if err == nil && err2 == nil {
		if strings.TrimSpace(ope) == "|" {
			res = strconv.FormatBool(left || right)
		} else if strings.TrimSpace(ope) == "+" {
			res = strconv.FormatBool(left && right)
		} else {
			res = strconv.FormatBool((left || right) && !(left && right))
		}
		return res
	} else {
		return ""
	}
}

func browseTree(t *Tree) string {
	gInsolvable++
	if gInsolvable > 10000 {
		return "false"
	}
	leftVal := ""
	rightVal := ""
	if strings.ContainsAny(t.Value, "+|^") {
		if t.Left != nil {
			leftVal = browseTree(t.Left)
		}
		if t.Right != nil {
			rightVal = browseTree(t.Right)
		}
		t.Value = operatorSolver(leftVal, t.Value, rightVal)
		t.Left = nil
		t.Right = nil
		return t.Value
	} else {
		if strings.ContainsAny(t.Value, "!") {
			val, err := strconv.ParseBool(vars[t.Value[1:]])
			if err == nil {
				return strconv.FormatBool(!val)
			} else {
				return "true"
			}
		}
		if vars[t.Value] == "" {
			newValue := backtraking(t.Value)
			return newValue
		}
		return vars[t.Value]
	}
}

func browseConclusionTree(t *Tree, res string) {
	if strings.ContainsAny(t.Value, "+") {
		if t.Left != nil {
			browseConclusionTree(t.Left, res)
		}
		if t.Right != nil {
			browseConclusionTree(t.Right, res)
		}
		return
	} else if res == "true" {
		if strings.ContainsAny(t.Value, "!") {
			val, err := strconv.ParseBool(res)
			if err == nil {
				strVal := strconv.FormatBool(!val)
				if vars[t.Value[1:]] != strVal && vars[t.Value[1:]] != "" {
					printErrorMsg("Contradiction in the variable values, please check the input!")
				}
				vars[t.Value[1:]] = strVal
			}
			return
		}
		if vars[t.Value] != res && vars[t.Value] != "" {
			printErrorMsg("Contradiction in the variable values, please check the input!")
		}
		vars[t.Value] = res
		return
	}
	return
}

func checkLetterInConc(letter string, rules []sRule) []int {
	var tab []int
	for i := 0; i < len(rules); i++ {
		for j := 0; j < len(rules[i].Conclusion); j++ {
			if strings.ContainsAny(rules[i].Conclusion[j], letter) {
				tab = append(tab, i)
			}
		}
	}
	return tab
}

func resolve(file sFile) {
	lenQueries := len(file.Queries)
	gRules = file.Rules
	for i := 0; i < lenQueries; i++ {
		res := checkLetterInConc(file.Queries[i], file.Rules)
		for j := 0; j < len(res); j++ {
			tFact := fillTree(file.Rules[res[j]].Facts)
			querieRes := browseTree(tFact)
			browseConclusionTree(fillTree(file.Rules[res[j]].Conclusion), querieRes)
		}
	}
	return
}
