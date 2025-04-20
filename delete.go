package tsbuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ TdEngineSQLBuilder = (*DeleteBuilder)(nil)

type DeleteBuilder struct {
	from            string
	whereConditions []string
}

func NewDeleteBuilder() *DeleteBuilder {
	return &DeleteBuilder{
		whereConditions: make([]string, 0),
	}
}

func (s *DeleteBuilder) From(from string) *DeleteBuilder {
	s.from = from
	return s
}

func (s *DeleteBuilder) Where(conditions ...string) *DeleteBuilder {
	s.whereConditions = append(s.whereConditions, conditions...)
	return s
}

func (s *DeleteBuilder) Build() (string, error) {
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

func (s *DeleteBuilder) validate() error {
	if s.from == "" {
		return ErrEmptyTableName
	}

	return nil
}
