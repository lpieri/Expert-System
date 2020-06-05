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
func rulesLoop(querie string, rules []sRule) {
	for i := 0; i < len(rules); i++ {
		fmt.Printf("Vars = %#v\n", vars)
		/*-----------------------------------               CAS N*1 : A = B               ------------------------------------*/
		if len(rules[i].Facts) == 1 && len(rules[i].Conclusion) == 1 {
			negation := false
			if strings.Contains(rules[i].Facts[0], "!") || strings.Contains(rules[i].Conclusion[0], "!") {
				negation = true
			}
			if vars[rules[i].Facts[0]] != "" && vars[rules[i].Conclusion[0]] != "" { // si existent deja toute les 2
				//fmt.Println("Toutes les 2 pas vide = ", vars[rules[i].Facts[0]], "     |      ", vars[rules[i].Conclusion[0]])
				if vars[rules[i].Facts[0]] == vars[rules[i].Conclusion[0]] { //on verifie qu'elles sont egales
					continue
				} else { //sinon on print err contradictoire
					printErrorMsg("System as contradictions please review input")
				}
			} else { // sinon on assigne la valeur de celle qui esst dans fact a l'autre
				if vars[rules[i].Facts[0]] != "" {
					bFacts, err2 := strconv.ParseBool(vars[rules[i].Facts[0]])
					if err2 == nil {
						//fmt.Println("fact pas vide =", vars[rules[i].Facts[0]])
						if negation {
							vars[rules[i].Conclusion[0]] = strconv.FormatBool(!bFacts)
						} else {
							vars[rules[i].Conclusion[0]] = vars[rules[i].Facts[0]]
						}
					} else {
						printError(err2)
					}
				}
			}
		}
		/*-----------------------------------               CAS Aute               ------------------------------------*/
		
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
			rulesLoop(file.Queries[i], file.Rules)
		}
	}
	fmt.Println("Init = ", file.Init[:])
	fmt.Printf("Vars = %#v\n", vars)
	return
}
