package tsbuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ tdEngineSqlBuilder = (*dropTableBuilder)(nil)

type dropTableBuilder struct {
	tables []string
}

func NewDropTableBuilder() *dropTableBuilder {
	return &dropTableBuilder{
		tables: make([]string, 0),
	}
}

func (s *dropTableBuilder) Tables(tables ...string) *dropTableBuilder {
	s.tables = append(s.tables, tables...)
	return s
}

func (s *dropTableBuilder) Build() (string, error) {
	if err := s.validate(); err != nil {
		return "", fmt.Errorf("validate error: %w", err)
	}

	b := bytes.NewBuffer([]byte{})
	b.WriteString("DROP TABLE ")

	for idx, table := range s.tables {
		s.tables[idx] = "IF EXISTS " + table
	}

	b.WriteString(strings.Join(s.tables, ", "))

	b.WriteString(";")

	return b.String(), nil
}

func (s *dropTableBuilder) validate() error {
	if len(s.tables) == 0 {
		return fmt.Errorf("tables are required")
	}

	return nil
}
