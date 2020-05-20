package main

import "fmt"

/* idées pour resolve :
- prendre la premiere querie
- iterer sur les regles et essayer de resoudre celle qui contienne la lettre recherché (j'ai verifié c'est forcement UNE seule lettre les variables en fait)
- si pas resolu alors chercher une regle avec la lettre qui nous manque etc etc jusqu'a avoir essayé toutes les regles
- si apres avoir essayer toutes les regles on a trouvé alors ok on passe à la querie suivante,
- sinon alors on ecrit pas resolvable
*/

func rulesLoop(querie string, rules []sRule) {
	//for i := 0; i < len(rules); i++ {
	for i := 0; i < 1; i++ {
		fmt.Println(rules[i])
		//* fonction resolve ONE rule */
		//on remplace les variables
		fmt.Println("ces faits = ", rules[i].Facts)
		fmt.Println("impliquent= ", rules[i].Conclusion)
		fmt.Println("len faits= ", len(rules[i].Facts))
		fmt.Println("len Conclusion= ", len(rules[i].Conclusion))
		if len(rules[i].Facts) == 1 && len(rules[i].Conclusion) == 1 {
			
			
			// valeur, ok := vars[rules[i].Facts[0]]
			// valeur2, ok2 := vars[rules[i].Conclusion[0]]

			/*test si presence not*/
			// if ok && ok2 { // si existent deja toute les 2 on verifie qu'elles sont egales
			// 	if valeur == valeur2 {
			// 		continue
			// 	} else { //sinon on print err contradictoire
			// 		printErrorMsg("System as contradictions please review input")
			// 	}
			// } else {

			// }
		}
		//on rajoute les valeurs trouvées et on return ok
		//ou on ne trouve rien et on retrurn false
	}
}

func resolve(file sFile) {
	fmt.Println("The queries is", file.Queries[:], "?")
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
