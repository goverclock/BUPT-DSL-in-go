package DSL

import (
	"bufio"
	"log"
	"os"
)

type symbolType int
const (
	varName symbolType = iota
	blockName
	funcName
	keyword
)

type varType int
const (
	string_ varType = iota
	int_
	float_
)
func getVarType(s string) varType {
	switch s {
	case "string":
		return string_
	case "int":
		return int_
	case "float":
		return float_
	}
	log.Fatal("no such variable type:", s)
	return -1
}

type variable struct {
	val string
	valType varType
}

type block struct {
	name string
	statements []string	
}

type position struct {
	blockName string
	statementIndex int
}
type Script struct {
	rawLines  []string
	symbols   map[string]symbolType
	blocks    map[string]block
	variables map[string]variable
	funcs     map[string]func(*Script, []string)

	pos position
}

// load script from a file
func NewScript(fname string) *Script {
	// create object
	obj := new(Script)
	obj.blocks = make(map[string]block)
	obj.symbols = make(map[string]symbolType)
	obj.variables = make(map[string]variable)
	obj.funcs = make(map[string]func(*Script, []string))

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		obj.rawLines = append(obj.rawLines, scan.Text())
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// analysis the script
	obj.pos.blockName = "begin"
	obj.parse()

	return obj
}
