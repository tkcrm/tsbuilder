package tdenginefuncs

type asin struct {
	expr string
}

func Asin(expr string) TDEngineFunc {
	return &asin{expr}
}

func (s asin) String() string {
	return "ASIN(" + s.expr + ")"
}
