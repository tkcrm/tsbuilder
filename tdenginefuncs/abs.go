package tdenginefuncs

type abs struct {
	expr string
}

func Abs(expr string) TDEngineFunc {
	return &abs{expr}
}

func (s abs) String() string {
	return "ABS(" + s.expr + ")"
}
