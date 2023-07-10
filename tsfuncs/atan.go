package tsfuncs

type atan struct {
	expr string
}

func Atan(expr string) TDEngineFunc {
	return &atan{expr}
}

func (s atan) String() string {
	return "ATAN(" + s.expr + ")"
}
