package tsbuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ TdEngineSQLBuilder = (*STableBuilder)(nil)

type STableBuilder struct {
	name        string
	definitions []string
	options     []string
	tags        map[string]any
}

func NewSTableBuilder() *STableBuilder {
	return &STableBuilder{
		definitions: make([]string, 0),
		options:     make([]string, 0),
		tags:        make(map[string]any),
	}
}

func (s *STableBuilder) Name(name string) *STableBuilder {
	s.name = name
	return s
}

func (s *STableBuilder) Definitions(items ...string) *STableBuilder {
	s.definitions = append(s.definitions, items...)
	return s
}

func (s *STableBuilder) Options(items ...string) *STableBuilder {
	s.options = append(s.options, items...)
	return s
}

func (s *STableBuilder) Tags(tags map[string]any) *STableBuilder {
	if tags == nil {
		tags = make(map[string]any)
	}
	s.tags = tags
	return s
}

func (s *STableBuilder) Build() (string, error) {
	if err := s.validate(); err != nil {
		return "", fmt.Errorf("validate error: %w", err)
	}

	b := bytes.NewBuffer([]byte{})
	b.WriteString("CREATE STABLE IF NOT EXISTS ")

	// add name
	b.WriteString(s.name + " ")

	// add definitions
	if len(s.definitions) > 0 {
		b.WriteString("(" + strings.Join(s.definitions, ", ") + ") ")
	}

	// add tags
	if len(s.tags) > 0 {
		tagsKeys := make([]string, 0, len(s.tags))
		tagsValues := make([]any, 0, len(s.tags))
		for key, value := range s.tags {
			tagsKeys = append(tagsKeys, key)
			tagsValues = append(tagsValues, value)
		}

		b.WriteString("TAGS (")
		for index, value := range tagsValues {
			v, err := castType(value)
			if err != nil {
				return "", err
			}
			b.WriteString(tagsKeys[index] + " ")
			b.WriteString(v)

			if index != len(tagsValues)-1 {
				b.WriteString(", ")
			}
		}
		b.WriteString(")")
	}

	// add options
	if len(s.options) > 0 {
		b.WriteString(" ")
		b.WriteString(strings.Join(s.options, " "))
	}

	b.WriteString(";")

	return b.String(), nil
}

func (s STableBuilder) validate() error {
	if s.name == "" {
		return ErrEmptySTableName
	}

	return nil
}
