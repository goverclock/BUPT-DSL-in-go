package DSL

import "log"

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
