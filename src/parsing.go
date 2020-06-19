package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func removeComment(line string) string {
	tab := strings.Split(line, "#")
	trim := strings.TrimSpace(tab[0])
	newLine := string(trim)
	return newLine
}

func addVar(tab []string) {
	var lettre string
	for i := 0; i < len(tab); i++ {
		lettre = tab[i]
		if tab[i] != "+" && tab[i] != "^" && tab[i] != "|" {
			if strings.Contains(tab[i], "!") {
				lettre = strings.Replace(tab[i], "!", "", -1)
			}
			if strings.Contains(tab[i], "(") {
				lettre = strings.Replace(tab[i], "(", "", -1)
			}
			if strings.Contains(tab[i], ")") {
				lettre = strings.Replace(tab[i], ")", "", -1)
			}
			vars[lettre] = ""
		}
	}
}

func checkParenthese(line string) bool {
	nbPtOpen := 0
	nbPtClosed := 0
	s := strings.ReplaceAll(line, " ", "")
	len := len(s)
	for i := 0; i < len; i++ {
		if s[i] == '(' {
			if i > 0 {
				if s[i-1] != '|' && s[i-1] != '+' && s[i-1] != '^' && s[i-1] != '(' {
					return false
				}
			} else if s[i+1] == '|' || s[i+1] == '+' || s[i+1] == '^' {
				return false
			}
			nbPtOpen++
		}
		if s[i] == ')' {
			if s[i-1] == '|' || s[i-1] == '+' || s[i-1] == '^' {
				return false
			} else if i+1 < len {
				if s[i+1] != '|' && s[i+1] != '+' && s[i+1] != '^' && s[i+1] != ')' {
					return false
				}
			}
			nbPtClosed++
		}
	}
	if nbPtOpen != nbPtClosed {
		return false
	}
	return true
}

func checkError(lineSplit []string) bool {
	for i := 0; i < len(lineSplit); i++ {
		if checkParenthese(lineSplit[i]) == false {
			return false
		}
		p := strings.Split(strings.TrimSpace(lineSplit[i]), "(")
		y := 0
		if p[0] == "" {
			y++
		}
		for ; y < len(p); y++ {
			s := strings.Split(strings.TrimSpace(p[y]), ")")[0]
			if s == "" {
				continue
			}
			if s[len(s)-1] == '+' || s[len(s)-1] == '|' || s[len(s)-1] == '^' {
				s = s[:len(s)-1]
			}
			s = strings.TrimSpace(s)
			re := regexp.MustCompile("^\\!?[A-Z]{1}(\\s+(\\+|\\||\\^)?\\s+\\!?[A-Z]{1})*")
			rest := re.Split(s, -1)
			for r := 0; r < len(rest); r++ {
				if rest[r] != "" {
					return false
				}
			}
		}
	}
	return true
}

func getRule(line string) sRule {
	if strings.Contains(line, "<=>") {
		printErrorMsg("'<=>' This is a bonus")
	} else if strings.Contains(line, "=>") == false {
		printErrorMsg("The file is badly formatted, please check it.")
	}
	lineSplit := strings.Split(line, "=>")
	if checkError(lineSplit) == false {
		printErrorMsg("The file is badly formatted, please check it.")
	}
	facts := strings.Split(strings.TrimSpace(lineSplit[0]), " ")
	conclusion := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
	if len(lineSplit) != 2 {
		printErrorMsg("Error no '[fact] => [conclusion]' found in rules, please review the format of the input file")
	}
	rule := sRule{Conclusion: conclusion, Facts: facts}
	addVar(facts)
	addVar(conclusion)
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
		if file.Queries == nil || file.Rules == nil {
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
