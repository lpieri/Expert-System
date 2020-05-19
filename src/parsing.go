package main

import (
	"io/ioutil"
	"strings"
)

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

func getRule(line string) Rule {
	if strings.Contains(line, "<=>") {
		printErrorMsg("'<=>' This is a bonus")
	}
	lineSplit := strings.Split(line, "=>")
	facts := strings.Split(strings.TrimSpace(lineSplit[0]), " ")
	conclusion := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
	rule := Rule{Conclusion: conclusion, Facts: facts}
	return rule
}

func openFile(fileName string) []Rule {
	var rules []Rule
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
			} else {
				continue
			}
		}
	}
	return rules
}

func checkArgs(arguments []string) {
	if len(arguments) < 3 {
		printUsage()
	} else {
		for i := 1; i < len(arguments); i++ {
			if arguments[i] != "--input" && arguments[i] != "-i" {
				printUsage()
			} else {
				openFile(arguments[i+1])
				i++
			}
		}
	}
}
