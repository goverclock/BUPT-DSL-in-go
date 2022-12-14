package DSL

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// system functions
// implement here and add to executor.go/initFuncs()

func save(s *Script, args []string) {
	lhv := s.variables[args[0]]
	rhv := args[1]
	t, ok := s.variables[rhv]
	if ok {	// rhv is a variable
		rhv = t.val
	} else if lhv.valType == string_ {	// rhv is a string
		rhv = rhv[1 : len(args[1])-1]
	}	// else rhv is a number
	lhv.val = rhv
	s.variables[args[0]] = lhv
	s.finish(position{})
}

func say(s *Script, args []string) {
	sentence := args[0][1 : len(args[0])-1] // remote quote
	// replace variable name to value
	for {
		pos := strings.Index(sentence, "${") // find variable ${var_name}
		if pos < 0 {
			break
		}
		ed := strings.Index(sentence, "}")
		toReplace := sentence[pos : ed+1] // replace with value
		varVal := s.variables[toReplace[2:len(toReplace)-1]].val
		sentence = strings.Replace(sentence, toReplace, varVal, 1)
	}

	sentence = "客服:" + sentence
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

func add(s *Script, args []string) {
	lname := args[0]
	rname := args[1]
	lvar := s.variables[lname]
	rvar := s.variables[rname]

	// convert to float if they are
	lv := 0.0
	rv := 0.0
	item := lvar
	if lvar.valType == float_ || lvar.valType == int_ || rvar.valType == float_ || rvar.valType == int_ {
		var err error
		lv, err = strconv.ParseFloat(lvar.val, 64)
		if err != nil {
			log.Fatal(err)
		}
		rv, err = strconv.ParseFloat(rvar.val, 64)
		if err != nil {
			log.Fatal(err)
		}
		item.val = strconv.FormatFloat(lv+rv, 'f', -1, 64)
	} else {
		item.val = lvar.val + rvar.val
	}
	s.variables[lname] = item

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
