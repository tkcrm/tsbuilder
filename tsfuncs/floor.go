package tsfuncs

type floor struct {
	expr string
}

func Floor(expr string) TDEngineFunc {
	return &floor{expr}
}

func (s floor) String() string {
	return "FLOOR(" + s.expr + ")"
}
