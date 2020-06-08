package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* idées pour resolve :
Arbre binaire
- get_minus_prio() ==> trouve moins prioritaire
- get_parentheses() ==> trouve les panrenthese et leur donne priorite ou non (a appaler dans get_minus_prio)
- construire arbre binaire
- resoudre en parcourant arbre
*/

func fillTree(querie string, rule sRule) *Tree {
	fmt.Printf("Vars = %#v\n", vars)
	if len(rule.Facts) == 1 && len(rule.Conclusion) == 1 {
		negation := false
		if strings.Contains(rule.Facts[0], "!") || strings.Contains(rule.Conclusion[0], "!") {
			negation = true
		}
		if vars[rule.Facts[0]] != "" && vars[rule.Conclusion[0]] != "" { // si existent deja toute les 2
			if vars[rule.Facts[0]] != vars[rule.Conclusion[0]] { //on verifie qu'elles sont egales
				printErrorMsg("System as contradictions please review input")
			}
		} else { // sinon on assigne la valeur de celle qui esst dans fact a l'autre
			if vars[rule.Facts[0]] != "" {
				bFacts, err2 := strconv.ParseBool(vars[rule.Facts[0]])
				if err2 == nil {
					if negation {
						vars[rule.Conclusion[0]] = strconv.FormatBool(!bFacts)
					} else {
						vars[rule.Conclusion[0]] = vars[rule.Facts[0]]
					}
				} else {
					printError(err2)
				}
			}
		}
		return nil
	} else {
		titi := newTree(rule.Facts)
		println("\nTREEEEEE : ", treeToString(titi))
		return titi
	}
}

func browseTree(t *Tree) string {
	if t.Left != nil {
		println("t left = ", t.Left.Value)
		leftVal := browseTree(t.Left)
	}
	if t.Right != nil {
		println("t right = ", t.Right.Value)
		rightVal := browseTree(t.Right)
	}

	// 	println("t value = ", fmt.Sprint(t.Value))
	// 	if vars[t.Left.Value] != "" && vars[t.Right.Value] != "" {
	// 		if t.Value == "|" {
	// 			left, err := strconv.ParseBool(vars[t.Left.Value])
	// 			right, err2 := strconv.ParseBool(vars[t.Right.Value])
	// 			if err == nil && err2 == nil {
	// 				res := left || right
	// 				print("res = ", res)
	// 				return strconv.FormatBool(res)
	// 			}
	// 		}
	// 	} else {
	// 		return ""
	// 	}
}

func resolve(file sFile) {
	fmt.Println("The queries is", file.Queries[:], "?")
	fmt.Println(file.Rules)
	lenQueries := len(file.Queries)
	for i := 0; i < lenQueries; i++ {
		if vars[string(file.Queries[i])] != "" {
			fmt.Println(file.Queries[i], "is", vars[string(file.Queries[i])])
			file.Queries = removeIndexFormTab(i, file.Queries)
			i--
			lenQueries--
			continue
		} else {
			cmpt := 0
			for j := 0; j < len(file.Rules); j++ {
				cmpt++
				t := fillTree(file.Queries[i], file.Rules[j])
				if t == nil {
					continue
				} else if cmpt > len(file.Rules) {
					printErrorMsg("resolution impossible, please change the input")
					break
				}
				browseTree(t)
			}
		}
	}
	fmt.Println("Init = ", file.Init[:])
	fmt.Printf("Vars = %#v\n", vars)
	return
}
