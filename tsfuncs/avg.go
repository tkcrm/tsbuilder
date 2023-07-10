package tsfuncs

type avg struct {
	expr string
}

func Avg(expr string) TDEngineFunc {
	return &avg{expr}
}

func (s avg) String() string {
	return "AVG(" + s.expr + ")"
}
