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
	for i := 0; i < len(rules); i++ {
		fmt.Println(rules[i])
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
