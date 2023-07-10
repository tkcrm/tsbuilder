package tdenginefuncs

type cos struct {
	expr string
}

func Cos(expr string) TDEngineFunc {
	return &cos{expr}
}

func (s cos) String() string {
	return "COS(" + s.expr + ")"
}
