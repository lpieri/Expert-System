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
	// println("Enter in backtraking -- Letter is:", newLetter)
	tab := checkLetterInConc(newLetter, gRules)
	for i := 0; i < len(tab); i++ {
		newFact := fillTree(gRules[tab[i]].Facts)
		// println("Backtraking tree facts:", treeToString(newFact))
		newValue := browseTree(newFact)
		if newValue == "" {
			/*		if (leftVal == "true" || rightVal == "true") && ope == "|" {
					return "true"
					} else if (leftVal == "false" || rightVal == "false") && ope == "+" {
						return "false"
					} else if (leftVal == "true" || rightVal == "true") && ope == "^" {
						return "true"
					}*/
			return "false"
		}
		if newValue == "true" {
			// println("newLetter:", newLetter, "=", newValue)
			vars[newLetter] = newValue
		}
		return newValue
	}
	return "false"
}

func operatorResolver(leftVal string, ope string, rightVal string) string {
	// println("Enter dans operatorResolver avec :", leftVal, "et", rightVal)
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
		// println("Exit de operatorResolver 1 et res = ", res, "\n")
		return res
	} else {
		// println("Exit de operatorResolver 2 et res = ", res, "\n")
		return ""
	}

}

func browseTree(t *Tree) string {
	leftVal := ""
	rightVal := ""
	// println("Enter in browseTree with tree:", treeToString(t))
	if strings.ContainsAny(t.Value, "+|^") {
		if t.Left != nil {
			leftVal = browseTree(t.Left)
		}
		if t.Right != nil {
			rightVal = browseTree(t.Right)
		}
		// println("Tree before the before:", treeToString(t))
		// println("Before -- leftVal:", leftVal, "rightVal:", rightVal)
		// if leftVal == "" {
		// 	leftVal = backtraking(t.Left.Value)
		// }
		// if rightVal == "" {
		// 	rightVal = backtraking(t.Right.Value)
		// }
		// println("After -- leftVal:", t.Left.Value, "=", leftVal, "rightVal:", t.Right.Value, "=", rightVal, "\n")
		// BACKTRAKING --- ici ---
		// avec test est-ce que lettre inconnue est dans une partie droite? (if ==> [contient LETTRE])
		// if backtraking est false alors on fait les tests suivants :
		t.Value = operatorResolver(leftVal, t.Value, rightVal)
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
			// println("Exit browseTree t.Value: ", t.Value, "=", newValue)
			return newValue
		}
		// println("Exit browseTree t.Value: ", t.Value, "=", vars[t.Value])
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
		//for chaque regles
		for j := 0; j < len(res); j++ {
			tFact := fillTree(file.Rules[res[j]].Facts)
			// println("before browseTree of tree:", treeToString(tFact))
			querieRes := browseTree(tFact)
			// println("result of tree:", treeToString(tFact), "=", querieRes, "\n")
			// fmt.Printf("vars = %v\n", vars)
			// println("conclusion tree:", treeToString(fillTree(file.Rules[res[j]].Conclusion)), "=", querieRes, "\n")
			browseConclusionTree(fillTree(file.Rules[res[j]].Conclusion), querieRes)
		}
		// println("int res checkletterinconc:", res, "facts:", treeToString(tFact))
		// tu essaye de la resoudre
		// si undefine tu check les tree et tu rappele cette function avec les du tree (ou le faire dans le tree)
		// Si queries toutes trouvée alors stop
		// res := browseTree(tFacts)
		// // println("res =", res)
		// browseConclusionTree(tConclusion, res)
	}
	return
}
