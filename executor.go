package DSL

import (
	"log"
)

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

func (s *Script) Run() {
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
