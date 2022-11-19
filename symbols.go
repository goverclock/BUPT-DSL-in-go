package DSL

type symbolType int

const (
	varname symbolType = iota
	blockname
	funcname
	keyword
)

type symbol struct {
	val     string
	valType symbolType
}

var syms map[string]symbol

func getSymbolType(s string) (t symbolType, ok bool) {
	if k, ok := syms[s]; ok {
		return k.valType, true
	}
	return 0, false
}

