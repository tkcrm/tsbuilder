package tdenginefuncs

type sum struct {
	expr string
}

func Sum(expr string) TDEngineFunc {
	return &sum{expr}
}

func (s sum) String() string {
	return "SUM(" + s.expr + ")"
}
