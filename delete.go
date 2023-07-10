package tdenginebuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ tdenginebuilder = (*deleteBuilder)(nil)

type deleteBuilder struct {
	from            string
	whereConditions []string
}

func NewDeleteBuilder() *deleteBuilder {
	return &deleteBuilder{
		whereConditions: make([]string, 0),
	}
}

func (s *deleteBuilder) From(from string) *deleteBuilder {
	s.from = from
	return s
}

func (s *deleteBuilder) Where(conditions ...string) *deleteBuilder {
	s.whereConditions = append(s.whereConditions, conditions...)
	return s
}

func (s *deleteBuilder) Build() (string, error) {
	if err := s.validate(); err != nil {
		return "", fmt.Errorf("validate error: %w", err)
	}

	b := bytes.NewBuffer([]byte{})
	b.WriteString("DELETE FROM ")

	// add from
	b.WriteString(s.from + " ")

	// add where conditions
	if len(s.whereConditions) > 0 {
		b.WriteString("WHERE ")
		b.WriteString(strings.Join(s.whereConditions, " AND "))
	}

	b.WriteString(";")

	return b.String(), nil
}

func (s *deleteBuilder) validate() error {
	if s.from == "" {
		return ErrEmptyTableName
	}

	return nil
}
