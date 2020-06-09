package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func checkParenthese(line string) bool {
	nbPtOpen := 0
	// nbPtClosed := 0
	println("prout:", line)
	s := strings.ReplaceAll(line, " ", "")
	for i := 0; i < len(line); i++ {
		if s[i] == '(' {
			//si char avant( n'est pas dans "|^+("  (ou debur de chaine) ==> alors erreur
			//sinon ok, on continu
			// si char apres ( est dans "|^+" alors erreur  ==> erreur
			//sinon ok on continu
			nbPtOpen++
		}
		//si char apres ) n'est pas dans "|^+)" (ou fin de chaine) ==> alors erreur
		//sinon ok, on continu
		// si char avant ) est dans "|^+" alors erreur  ==> erreur
		//sinon ok on continu
	}
	//verif que nbPtOpen et nbPtClosed sont egaux sinon erreur
	// if ce n'est pas bon return false
	return true
}

func checkerror(lineSplit []string) bool {
	//decouper l
	//check apres decoupage
	fmt.Println(lineSplit)
	cmp := 0
	// var lastOpened int = 0
	for i := 0; i < len(lineSplit); i++ {
		checkParenthese(lineSplit[i])
		s := lineSplit[i]
		for j := 0; j < len(s); j++ {
			if s[i] == '(' {
				// lastOpened = i
				cmp++
			} else if s[i] == ')' {
				//sub string lastopened : i
				// checkerror(substring)
				//resolve(substring) return TRUE or FALSE OR undef
				cmp--
				//lastOpened = i
			}
		}
		//si cmp != 0 ==> error
	}

	//A + ((B + C) + D)

	// for i := 0; i < len(lineSplit); i++ {
	// 	p := strings.Split(strings.TrimSpace(lineSplit[i]), "(")
	// 	fmt.Println(p)
	// 	var i int
	// 	if p[0] == "!" || p[0] == "" {
	// 		i = 1
	// 	} else {
	// 		i = 0
	// 	}
	// 	for ; i < len(p); i++ {
	// 		s := strings.Split(strings.TrimSpace(p[i]), ")")[0]
	// 		fmt.Println(s)
	// 		re := regexp.MustCompile("^\\!?[A-Z]{1}(\\s+(\\+|\\||\\^)?\\s+\\!?[A-Z]{1})*")
	// 		fmt.Println(re.MatchString(s))
	// 	}

	// (H + D (H + ) P)
	// (H + D (H + P))
	//A + (B | C) + C => A + , B | C) + C
	/*
		Normal:
		[!(H ^ G)   !(F + C)]
		[! H ^ G)]
		H ^ G
		true
		[! F + C)]
		F + C
		true

		Pas normal: valide h + espace c'est degueux
		[(H + D (H + ) P)   C]
		[ H + D  H + ) P)]
		H + D
		true
		H +   ???
		true
		[C]
		C
		true

		Normal:
		[(H + D (H + P))   C]
		[ H + D  H + P))]
		H + D
		true
		H + P
		true
		[C]
		C
		true
	*/
	return true
}

func getRule(line string) sRule {

	//parenthese ok :  ^(\!)?(?<parenthse>\()?[A-Z]{1}(\s+(\+|\||\^)?\s+)(\!?[A-Z]{1})*(?(parenthse)\)|\s*)
	// re := regexp.MustCompile("^\\!?[A-Z]{1}(\\s+(\\+|\\||\\^)?\\s+\\!?[A-Z]{1})*")

	if strings.Contains(line, "<=>") {
		printErrorMsg("'<=>' This is a bonus")
	} else if strings.Contains(line, "=>") == false {
		printErrorMsg("The file is badly formatted, please check it.")
	}
	lineSplit := strings.Split(line, "=>")
	if checkerror(lineSplit) == false {
		printErrorMsg("The file is badly formatted, please check it.")
	}
	facts := strings.Split(strings.TrimSpace(lineSplit[0]), " ")
	conclusion := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
	addVar(facts)
	addVar(conclusion)
	rule := sRule{Conclusion: conclusion, Facts: facts}
	return rule
}

func parseFile(data string) sFile {
	var file sFile
	lines := strings.Split(data, "\n")
	nbLines := len(lines)
	for i := 0; i < nbLines; i++ {
		line := removeComment(string(lines[i]))
		if len(line) > 0 {
			if line[0] != '=' && line[0] != '?' {
				file.Rules = append(file.Rules, getRule(line))
			} else if line[0] == '=' {
				line = line[1:]
				for j := 0; j < len(line); j++ {
					file.Init = append(file.Init, string(line[j]))
					vars[string(line[j])] = "true"
				}
			} else if line[0] == '?' {
				line = line[1:]
				for j := 0; j < len(line); j++ {
					file.Queries = append(file.Queries, string(line[j]))
				}
			}
		} else {
			continue
		}
	}
	return file
}

func openFile(fileName string) sFile {
	var file sFile
	bits, err := ioutil.ReadFile(fileName)
	if err != nil {
		printError(err)
	} else {
		data := string(bits)
		file = parseFile(data)
		if file.Queries == nil || file.Init == nil || file.Rules == nil {
			printErrorMsg("Rules, Facts or Querie missing, please review your input file.")
		}
	}
	return file
}

func checkArgs(arguments []string) sFile {
	var file sFile
	if len(arguments) < 3 {
		printUsage()
	} else {
		for i := 1; i < len(arguments); i++ {
			if arguments[i] != "--input" && arguments[i] != "-i" {
				printUsage()
			} else {
				file = openFile(arguments[i+1])
				i++
			}
		}
	}
	return file
}
