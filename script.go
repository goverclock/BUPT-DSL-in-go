package DSL

import (
	"bufio"
	"log"
	"os"
)

type Script struct {
	rawLines []string
	symbols map[string]symbolType
	blocks   map[string]block
	variables map[string]variable

	position      struct {
		blockName      string
		statementIndex int
	}
}

// load script from a file
func NewScript(fname string) *Script {
	// create object
	obj := new(Script)
	obj.blocks = make(map[string]block)
	obj.symbols = make(map[string]symbolType)
	obj.variables = map[string]variable{}

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
	obj.position.blockName = "begin"
	obj.parse()

	return obj
}
