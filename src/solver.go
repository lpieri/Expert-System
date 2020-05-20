package main

import "fmt"

func resolve(file sFile) {
	/* idées pour resolve :
	- prendre la premiere querie
	- iterer sur les regles et essayer de resoudre celle qui contienne la lettre recherché (j'ai verifié c'est forcement UNE seule lettre les variables en fait)
	- si pas resolu alors chercher une regle avec la lettre qui nous manque etc etc jusqu'a avoir essayé toutes les regles
	- si apres avoir essayer toutes les regles on a trouvé alors ok on passe à la querie suivante,
	- sinon alors on ecrit pas resolvable
	*/
	fmt.Println("Regles = ", file.Rules[:])
	fmt.Println("Init = ", file.Init[:])
	fmt.Println("Queries = ", file.Queries[:])
	fmt.Printf("Vars = %#v\n", vars)
	return
}
