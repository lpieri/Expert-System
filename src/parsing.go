package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*Variables Gloabales*/
var vars = map[string]string{}

type Rule struct {
	Facts      []string
	Conclusion []string
}

func removeComment(line string) string {
	tab := strings.Split(line, "#")
	trim := strings.TrimSpace(tab[0])
	newLine := string(trim)
	return newLine
}

func addVar(tab []string) {
	for i := 0; i < len(tab); i++ {
		if tab[i] != "+" && tab[i] != "^" && tab[i] != "|" {
			if strings.Contains(tab[i], "!") {
				tab[i] = tab[i][1:]
			}
			vars[tab[i]] = ""
		}
	}
}

func getRule(line string) Rule {
	if strings.Contains(line, "<=>") {
		printErrorMsg("'<=>' This is a bonus")
	}
	lineSplit := strings.Split(line, "=>")
	facts := strings.Split(strings.TrimSpace(lineSplit[0]), " ")
	conclusion := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
	//on rempli le tableau avec les variables initialisées a false
	addVar(facts)
	addVar(conclusion)
	rule := Rule{Conclusion: conclusion, Facts: facts}
	return rule
}

func openFile(fileName string) ([]Rule, []string, []string) {
	var rules []Rule
	var queries []string
	var init []string
	bits, err := ioutil.ReadFile(fileName)

	if err != nil {
		printError(err)
	} else {
		data := string(bits)
		lines := strings.Split(data, "\n")
		nbLines := len(lines)
		for i := 0; i < nbLines; i++ {
			line := removeComment(string(lines[i]))
			if len(line) > 0 {
				if line[0] != '=' && line[0] != '?' {
					rules = append(rules, getRule(line))
				}
				if line[0] == '=' {
					line = line[1:]
					for j := 0; j < len(line); j++ {
						//on rempli le tableau facts
						init = append(init, string(line[j]))
						vars[string(line[j])] = "true" //on met les faits initiaux a true
					}
				}
				if line[0] == '?' {
					line = line[1:]
					for j := 0; j < len(line); j++ {
						//on rempli le tableau queries
						queries = append(queries, string(line[j]))
					}
				}
			} else {
				continue
			}
		}
		//Check erreur
		if queries == nil || init == nil || rules == nil {
			printErrorMsg("Rules, Facts or Querie missing, please review your input file.")
		}
	}
	return rules, init, queries
}

func resolve(rules []Rule, init []string, queries []string) {
	/* idées pour resolve :
	- prendre la premiere querie
	- iterer sur les regles et essayer de resoudre celle qui contienne la lettre recherché (j'ai verifié c'est forcement UNE seule lettre les variables en fait)
	- si pas resolu alors chercher une regle avec la lettre qui nous manque etc etc jusqu'a avoir essayé toutes les regles
	- si apres avoir essayer toutes les regles on a trouvé alors ok on passe à la querie suivante,
	- sinon alors on ecrit pas resolvable
	*/
	fmt.Println("Regles = ", rules[:])
	fmt.Println("init = ", init[:])
	fmt.Println("Queries = ", queries[:])
	fmt.Printf("vars = %#v\n", vars)
	return
}

func checkArgs(arguments []string) {

	if len(arguments) < 3 {
		printUsage()
	} else {
		for i := 1; i < len(arguments); i++ {
			if arguments[i] != "--input" && arguments[i] != "-i" {
				printUsage()
			} else {
				rules, facts, queries := openFile(arguments[i+1])
				resolve(rules, facts, queries)
				i++
			}
		}
	}
}
