package tdenginefuncs

type round struct {
	expr string
}

func Round(expr string) TDEngineFunc {
	return &round{expr}
}

func (s round) String() string {
	return "ROUND(" + s.expr + ")"
}
