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
	str varType = iota
	integer 
	float
)
func getVarType(s string) varType {
	switch s {
	case "str":
		return str
	case "integer":
		return integer	
	case "float":
		return float
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
