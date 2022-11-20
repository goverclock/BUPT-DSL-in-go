package DSL

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func (s *Script) initFuncs() {
	s.funcs["say"] = say
	s.funcs["input"] = input
	s.funcs["_match"] = _match
	s.funcs["goto"] = _goto
}

func (s *Script) dispatchFunc(f string, args string) {
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

func parseStatement(s string) (name string, args string) {
	for i, v := range s {
		if v == '(' {
			return s[:i], s[i+1 : len(s)-1]
		}
	}
	return "", ""
}

func (s *Script) Run() {
	s.initFuncs()

	for {
		curBlock := s.pos.blockName
		curStaInd := s.pos.statementIndex
		// check end
		if len(s.blocks[curBlock].statements) == curStaInd {
			log.Println("finished, quitting")
			break
		}

		curSta := s.blocks[curBlock].statements[curStaInd]
		// function name & args
		funcName, funcArgs := parseStatement(curSta)

		// fmt.Println(funcName, funcArgs)
		s.dispatchFunc(funcName, funcArgs)
	}
}

// ******************** functions ********************** //

func say(s *Script, args string) {
	args = "客服:" + args[1:len(args)-1] // remove ""
	fmt.Println(args)

	s.finish(position{})
}

func input(s *Script, args string) {
	item := s.variables[args]
	// fmt.Println("before:", item)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1] // remove '\n'
	item.val = text
	s.variables[args] = item
	// fmt.Println("after:", item)

	s.finish(position{})
}

func _match(s *Script, args string) {
	as := strings.Split(args, ",")
	lhv := s.variables[as[0]].val
	rhv := strings.Clone(as[1])
	if rhv[0] == '"' && rhv[len(rhv)-1] == '"' {
		rhv = rhv[1 : len(rhv)-1]
	}
	if rhv == "default" || lhv == rhv {
		t := ""
		for i := 2; i < len(as); i++ {
			t += as[i]
		}
		s.dispatchFunc(parseStatement(t))
	} else {
		s.finish(position{})
	}
}

func _goto(s *Script, args string) {
	s.finish(position{args, 0})	
}
