package tsbuilder

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var _ TdEngineSQLBuilder = (*DropTableBuilder)(nil)

type DropTableBuilder struct {
	tables []string
}

func NewDropTableBuilder() *DropTableBuilder {
	return &DropTableBuilder{
		tables: make([]string, 0),
	}
}

func (s *DropTableBuilder) Tables(tables ...string) *DropTableBuilder {
	s.tables = append(s.tables, tables...)
	return s
}

func (s *DropTableBuilder) Build() (string, error) {
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

func (s *DropTableBuilder) validate() error {
	if len(s.tables) == 0 {
		return errors.New("tables are required")
	}

	return nil
}
