package tsfuncs

type binary struct {
	expr string
}

func Binary(expr string) TDEngineFunc {
	return &binary{expr}
}

func (s binary) String() string {
	return "BINARY(" + s.expr + ")"
}
