package tsbuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ TdEngineSQLBuilder = (*DatabaseBuilder)(nil)

type DatabaseBuilder struct {
	name    string
	options []string
}

func NewDatabaseBuilder() *DatabaseBuilder {
	return &DatabaseBuilder{
		options: make([]string, 0),
	}
}

func (s *DatabaseBuilder) Name(name string) *DatabaseBuilder {
	s.name = name
	return s
}

func (s *DatabaseBuilder) Options(options ...string) *DatabaseBuilder {
	s.options = append(s.options, options...)
	return s
}

func (s *DatabaseBuilder) Build() (string, error) {
	if err := s.validate(); err != nil {
		return "", fmt.Errorf("validate error: %w", err)
	}

	b := bytes.NewBuffer([]byte{})
	b.WriteString("CREATE DATABASE IF NOT EXISTS ")

	// add name
	b.WriteString(s.name + " ")

	// add options
	b.WriteString(strings.Join(s.options, " "))

	b.WriteString(";")

	return b.String(), nil
}

func (s DatabaseBuilder) validate() error {
	if s.name == "" {
		return ErrEmptyDatabaseName
	}

	return nil
}
