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

func logiqueresolv(val1 string, ope string, val2 string) string {
	res := ""
	println("dans logique resolved\nval1 = ", val1, "    ope = ", ope, "    val2 = ", val2)
	left, err := strconv.ParseBool(val1)
	right, err2 := strconv.ParseBool(val2)
	println("left = ", left, "  right = ", right)
	if err == nil && err2 == nil {
		println("left = ", left, "  right = ", right)
		if strings.TrimSpace(ope) == "|" {
			res = strconv.FormatBool(left || right)
			print("res = ", res)
			return res
		} else if strings.TrimSpace(ope) == "+" {
			res = strconv.FormatBool(left && right)
			print("res = ", res)
			return res
		} else {
			res = strconv.FormatBool((left || right) && !(left && right))
			print("res = ", res)
			return res
		}
	} else {
		fmt.Println("resolveNode: left:", left, "right:", right)
		if (left == true || right == true) && ope == "|" {
			return "true"
		} else if (left == false || right == false) && ope == "+" {
			return "false"
		}
		// fmt.Println("Return val of browseTree val: undefine")
		return "undefine"
	}
}

func browseTree(t *Tree) string {
	leftVal := ""
	rightVal := ""
	println("t value = ", fmt.Sprint(t.Value))
	//si val actuelle != lettre
	if strings.ContainsAny(t.Value, "+|^") {
		//on rappele browse tree : leftVal = browseTree(t.Left)
		if t.Left != nil {
			println("t left = ", t.Left.Value)
			leftVal = browseTree(t.Left)
		}
		if t.Right != nil {
			println("t right = ", t.Right.Value)
			rightVal = browseTree(t.Right)
		}
		t.Value = logiqueresolv(leftVal, t.Value, rightVal)
		t.Left = nil
		t.Right = nil
		return t.Value
	} else { //si lettre
		if strings.ContainsAny(t.Value, "!") {
			val, err := strconv.ParseBool(vars[t.Value[1:]])
			if err == nil {
				return strconv.FormatBool(!val)
			} else {
				printErrorMsg("ERROR")
				return "error"
			}
		}
		println("t.Value 2 = ", t.Value)
		println("vars[t.Value] = ", vars[t.Value])
		return vars[t.Value]
	}
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
				}
				res := browseTree(t)
				println("AFTER === ", treeToString(t), "   res = ", res)
				if res != "undefine" {
					println("lolilol")
					if vars[file.Queries[i]] != "" {
						println("lolilol222")
						if vars[file.Queries[i]] == res {
							println("lolilol3333")
							file.Queries = removeIndexFormTab(i, file.Queries)
							i--
							lenQueries--
						} else {
							printErrorMsg("Contradiction in the variable values, please check the input!")
						}
					} else {
						println("lolilol4444")
						vars[file.Queries[i]] = res
						file.Queries = removeIndexFormTab(i, file.Queries)
						i--
						lenQueries--
					}
				}
			}
		}
	}
	fmt.Println("Init = ", file.Init[:])
	fmt.Printf("Vars = %#v\n", vars)
	return
}
