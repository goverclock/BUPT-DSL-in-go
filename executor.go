package DSL

import (
	"fmt"
	"log"
)

func (s *Script) dispatchFunc(f string, args string) {
	if fn, ok := s.funcs[f]; ok {
		fn(s, args)
	} else {
		log.Fatal("no such function:", f)
	}
}

func (s *Script) initFuncs() {
	s.funcs["say"] = say
}

func (s *Script) Run() {
	s.initFuncs()

	for {
		curBlock := s.position.blockName
		curStaInd := s.position.statementIndex
		// s.funcs[]
		curSta := s.blocks[curBlock].statements[curStaInd]
		// function name & args
		funcName := ""
		funcArgs := ""
		for i, v := range curSta {
			if v == '(' {
				funcName = curSta[:i]
				funcArgs = curSta[i+1 : len(curSta)-1]
				break
			}
		}
		// fmt.Println(funcName, funcArgs)
		s.dispatchFunc(funcName, funcArgs)
	}
}

// ******************** functions ********************** //

func say(s *Script, args string) {
	fmt.Println("Called say():", args)
}
