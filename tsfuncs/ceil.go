package tsfuncs

type ceil struct {
	expr string
}

func Ceil(expr string) TDEngineFunc {
	return &ceil{expr}
}

func (s ceil) String() string {
	return "CEIL(" + s.expr + ")"
}
