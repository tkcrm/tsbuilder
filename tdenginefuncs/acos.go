package tdenginefuncs

type acos struct {
	expr string
}

func Acos(expr string) TDEngineFunc {
	return &acos{expr}
}

func (s acos) String() string {
	return "ACOS(" + s.expr + ")"
}
