package DSL

import (
	"fmt"
	"strings"
)

// initialize symbols
func (s *Script) parse() {
	inVar := false
	inBlock := false
	bName := ""
	bStatements := []string{}
	for _, line := range s.rawLines {
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

		// blocks
		if len(words) > 1 && words[1] == "{" {
			inBlock = true
			bName = words[0]
			continue
		}
		if inBlock {
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

	//

	showSymbols(s)
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
