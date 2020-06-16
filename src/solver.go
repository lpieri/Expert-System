package main

import "strings"

/* idées pour resolve :
Arbre binaire
- get_minus_prio() ==> trouve moins prioritaire
- get_parentheses() ==> trouve les panrenthese et leur donne priorite ou non (a appaler dans get_minus_prio)
- construire arbre binaire
- resoudre en parcourant arbre
*/

func fillTree(querie string, toTree []string) *Tree {
	tree := newTree(toTree)
	return tree
}

// func operatorResolver(leftVal string, ope string, rightVal string) string {
// 	res := ""
// 	left, err := strconv.ParseBool(leftVal)
// 	right, err2 := strconv.ParseBool(rightVal)
// 	if err == nil && err2 == nil {
// 		if strings.TrimSpace(ope) == "|" {
// 			res = strconv.FormatBool(left || right)
// 		} else if strings.TrimSpace(ope) == "+" {
// 			res = strconv.FormatBool(left && right)
// 		} else {
// 			res = strconv.FormatBool((left || right) && !(left && right))
// 		}
// 		return res
// 	} else {
// 		// BACKTRAKING --- ici ---
// 		// avec test est-ce que lettre inconnue est dans une partie droite? (if ==> [contient LETTRE])
// 		// if backtraking est false alors on fait les tests suivants :

// 		if (leftVal == "true" || rightVal == "true") && ope == "|" {
// 			return "true"
// 		} else if (leftVal == "false" || rightVal == "false") && ope == "+" {
// 			return "false"
// 		} else if (leftVal == "true" || rightVal == "true") && ope == "^" {
// 			return "true"
// 		}

// 		return "undefine"
// 	}
// }

// func browseTree(t *Tree) string {
// 	leftVal := ""
// 	rightVal := ""
// 	if strings.ContainsAny(t.Value, "+|^") {
// 		if t.Left != nil {
// 			leftVal = browseTree(t.Left)
// 		}
// 		if t.Right != nil {
// 			rightVal = browseTree(t.Right)
// 		}
// 		t.Value = operatorResolver(leftVal, t.Value, rightVal)
// 		t.Left = nil
// 		t.Right = nil
// 		return t.Value
// 	} else {
// 		if strings.ContainsAny(t.Value, "!") {
// 			val, err := strconv.ParseBool(vars[t.Value[1:]])
// 			if err == nil {
// 				return strconv.FormatBool(!val)
// 			} else {
// 				return "true"
// 			}
// 		}
// 		return vars[t.Value]
// 	}
// }

// func browseConclusionTree(t *Tree, res string) {
// 	if strings.ContainsAny(t.Value, "+") {
// 		if t.Left != nil {
// 			browseConclusionTree(t.Left, res)
// 		}
// 		if t.Right != nil {
// 			browseConclusionTree(t.Right, res)
// 		}
// 		return
// 	} else {
// 		if strings.ContainsAny(t.Value, "!") {
// 			val, err := strconv.ParseBool(res)
// 			if err == nil {
// 				strVal := strconv.FormatBool(!val)
// 				vars[t.Value[1:]] = strVal
// 			}
// 			return
// 		} else if res != "undefine" {
// 			if vars[t.Value] == "" {
// 				vars[t.Value] = res
// 				return
// 			} else if vars[t.Value] != res {
// 				printErrorMsg("Contradiction in the variable values, please check the input!")
// 			}
// 		}
// 	}
// 	return
// }

func checkLetterInConc(letter string, rules []sRule) int {
	for i := 0; i < len(rules); i++ {
		for j := 0; j < len(rules[i].Conclusion); j++ {
			if strings.ContainsAny(rules[i].Conclusion[j], letter) {
				return i
			}
		}
	}
	return -1
}

func resolve(file sFile) {
	lenQueries := len(file.Queries)
	for i := 0; i < lenQueries; i++ {
		checkLetterInConc(file.Queries[i], file.Rules)
		// check si ta queri est conclusion ...
		// si conclusion call fact
		// tu essaye de la resoudre
		// si undefine tu check les tree et tu rappele cette function avec les du tree (ou le faire dans le tree)
		// tFacts := fillTree(file.Queries[i], file.Rules[j].Facts)
		// // println("tFacts:", treeToString(tFacts))
		// // tConclusion := fillTree(file.Queries[i], file.Rules[j].Conclusion)
		// if tFacts == nil {
		// 	continue
		// }
		// Si queries toutes trouvée alors stop
		// res := browseTree(tFacts)
		// // println("res =", res)
		// browseConclusionTree(tConclusion, res)
	}
	return
}
