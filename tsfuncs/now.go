package tsfuncs

type now struct{}

func Now() TDEngineFunc {
	return &now{}
}

func (s now) String() string {
	return "NOW()"
}
