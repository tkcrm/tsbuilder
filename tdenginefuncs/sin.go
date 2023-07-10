package tdenginefuncs

type sin struct {
	expr string
}

func Sin(expr string) TDEngineFunc {
	return &sin{expr}
}

func (s sin) String() string {
	return "SIN(" + s.expr + ")"
}
