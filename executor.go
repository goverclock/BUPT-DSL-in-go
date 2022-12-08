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
	s.funcs["save"] = save
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

// TODO: return name string, args []string
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

		// fmt.Println(funcName, funcArgs)
		s.dispatchFunc(funcName, funcArgs)
	}
}

// ******************** functions ********************** //

func save(s *Script, args []string) {
	item := s.variables[args[0]]
	item.val = args[1]
	s.finish(position{})
}

func say(s *Script, args []string) {
	sentence := "客服:" + args[0][1:len(args[0])-1]
	fmt.Println(sentence)
	s.finish(position{})
}

func input(s *Script, args []string) {
	item := s.variables[args[0]]
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1] // remove '\n'
	item.val = text
	s.variables[args[0]] = item

	s.finish(position{})
}

func _match(s *Script, args []string) {
	lhv := s.variables[args[0]].val
	rhv := strings.Clone(args[1])
	// if rhv is a string
	if rhv[0] == '"' && rhv[len(rhv)-1] == '"' {
		rhv = rhv[1 : len(rhv)-1]
	}

	// if lhv matches rhv
	if rhv == "default" || lhv == rhv {
		t := ""
		for i := 2; i < len(args); i++ {
			t += args[i]
		}
		s.dispatchFunc(parseStatement(t))
	} else {
		s.finish(position{})
	}
}

func _goto(s *Script, args []string) {
	s.finish(position{args[0], 0})
}
