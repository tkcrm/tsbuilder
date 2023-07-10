package tdenginebuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ tdenginebuilder = (*databaseBuilder)(nil)

type databaseBuilder struct {
	name    string
	options []string
}

func NewDatabaseBuilder() *databaseBuilder {
	return &databaseBuilder{
		options: make([]string, 0),
	}
}

func (s *databaseBuilder) Name(name string) *databaseBuilder {
	s.name = name
	return s
}

func (s *databaseBuilder) Options(options ...string) *databaseBuilder {
	s.options = append(s.options, options...)
	return s
}

func (s *databaseBuilder) Build() (string, error) {
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

func (s databaseBuilder) validate() error {
	if s.name == "" {
		return ErrEmptyDatabaseName
	}

	return nil
}
