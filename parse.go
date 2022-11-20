package DSL

import (
	"fmt"
	"strings"
)

// initialize symbols
func (s *Script) parse() {
	inVar := false
	inBlock := false
	inSwitch := false
	switchArgs := []string{}
	bName := ""
	bStatements := []string{}
	for _, line := range s.rawLines {
		line = strings.TrimSpace(line)
		words := strings.Fields(line)
		// variable
		if line == "var (" {
			inVar = true
			continue
		}
		if inVar {
			if line == ")" {
				inVar = false
				continue
			}
			vname := words[0]
			vt := words[1]
			s.symbols[vname] = varName                        // save symbol
			s.variables[vname] = variable{"", getVarType(vt)} // save variable
		}

		// switches
		// match in "你好" goto hello
		if inBlock && !inSwitch && len(words) > 1 && words[0] == "switch" {
			inSwitch = true
			switchArgs = append(switchArgs, "match")
			switchArgs = append(switchArgs, words[1])
			continue
		}
		if inSwitch {
			if line == "}" {
				inSwitch = false
				switchArgs = nil
				continue
			}
			if line != "" {
				sta := switchArgs[0] + "(" + switchArgs[1]
				for _, w := range words {
					sta += "," + w
				}
				sta += ")"
				bStatements = append(bStatements, sta)
			}
		}

		// blocks
		if !inBlock && len(words) > 1 && words[1] == "{" {
			inBlock = true
			bName = words[0]
			continue
		}
		if inBlock && !inSwitch {
			if line == "}" {
				s.blocks[bName] = block{bName, bStatements}
				inBlock = false
				bName = ""
				bStatements = nil
				continue
			}
			if line != "" {
				bStatements = append(bStatements, line)
			}
		}
	}

	// showSymbols(s)
}

func showSymbols(s *Script) {
	fmt.Println("Symbols:\n", s.symbols)
	fmt.Println("Variables:\n", s.variables)
	fmt.Println("Blocks:")
	for _, b := range s.blocks {
		fmt.Println(b.name)
		for _, sta := range b.statements {
			fmt.Println(sta)
		}
	}
}
