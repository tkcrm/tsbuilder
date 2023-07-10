package tdenginefuncs

import "strings"

type count struct {
	columns []string
}

func Count(columns ...string) TDEngineFunc {
	return &count{columns}
}

func (s count) String() string {
	return "COUNT(" + strings.Join(s.columns, ", ") + ")"
}
