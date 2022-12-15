package DSL

import (
	"strings"
)

// load system functions from fns.go
func (s *Script) initFuncs() {
	s.funcs["say"] = say
	s.funcs["input"] = input
	s.funcs["_match"] = _match
	s.funcs["goto"] = _goto
	s.funcs["save"] = save
	s.funcs["add"] = add
}

// load user functions from user_fns.go
func (s *Script) initUserFuncs() {
	s.funcs["catfact"] = catfact
	s.funcs["dogfact"] = dogfact
}

// analysis the script
func (s *Script) parse() {
	s.initFuncs()
	s.initUserFuncs()

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
			s.variables[vname] = variable{"", getVarType(vt)} // save variable
		}

		// switches
		if inBlock && !inSwitch && len(words) > 1 && words[0] == "switch" {
			inSwitch = true
			switchArgs = append(switchArgs, "_match")
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
				sta += ")"		// sta == "_match(in,"你好",goto(hello))"
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
}
