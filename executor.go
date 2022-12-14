package DSL

import (
	"log"
)

// from fns.go
func (s *Script) initFuncs() {
	s.funcs["say"] = say
	s.funcs["input"] = input
	s.funcs["_match"] = _match
	s.funcs["goto"] = _goto
	s.funcs["save"] = save
	s.funcs["add"] = add
}

func (s *Script) initUserFuncs() {
	s.funcs["catfact"] = catfact
	s.funcs["dogfact"] = dogfact
}

func (s *Script) dispatchFunc(f string, args []string) {
	if fn, ok := s.funcs[f]; ok {
		fn(s, args)
	} else {
		log.Fatal("no such function:", f)
	}
}

// arg: next statement to execute
// {"", 0} for next ind
func (s *Script) finish(p position) {
	if p.blockName == "" {
		s.pos.statementIndex++
	} else {
		s.pos = p
	}
}

func parseStatement(s string) (name string, args []string) {
	for i, v := range s {
		if v == '(' {
			arg := s[i+1 : len(s)-1]
			return s[:i], parseArgs(arg)
		}
	}
	return "", nil
}

func parseArgs(s string) []string {
	inQuote := false
	ret := []string{}
	last := 0
	for i, v := range s {
		if v == '"' {
			inQuote = !inQuote
		}
		if inQuote {
			continue
		}
		if v == ',' {
			ret = append(ret, s[last:i])
			last = i + 1
		}
		if i == len(s)-1 {
			ret = append(ret, s[last:])
		}
	}
	return ret
}

func (s *Script) Run() {
	s.initFuncs()
	s.initUserFuncs()

	for {
		curBlock := s.pos.blockName
		curStaInd := s.pos.statementIndex
		// check end
		if len(s.blocks[curBlock].statements) == curStaInd {
			log.Println("script finished, quitting")
			break
		}

		curSta := s.blocks[curBlock].statements[curStaInd]
		// function name & args
		funcName, funcArgs := parseStatement(curSta)

		s.dispatchFunc(funcName, funcArgs)
	}
}
