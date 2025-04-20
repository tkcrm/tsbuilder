package tsbuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ TdEngineSQLBuilder = (*CreateTableBuilder)(nil)

type CreateTableBuilder struct {
	tableName  string
	sTableName string
	tags       map[string]any
}

func NewCreateTableBuilder() *CreateTableBuilder {
	return &CreateTableBuilder{
		tags: make(map[string]any),
	}
}

func (s *CreateTableBuilder) TableName(tableName string) *CreateTableBuilder {
	s.tableName = tableName
	return s
}

func (s *CreateTableBuilder) STable(sTableName string) *CreateTableBuilder {
	s.sTableName = sTableName
	return s
}

func (s *CreateTableBuilder) Tags(tags map[string]any) *CreateTableBuilder {
	if tags == nil {
		tags = make(map[string]any)
	}
	s.tags = tags
	return s
}

func (s *CreateTableBuilder) Build() (string, error) {
	if err := s.validate(); err != nil {
		return "", fmt.Errorf("validate error: %w", err)
	}

	b := bytes.NewBuffer([]byte{})
	b.WriteString("CREATE TABLE IF NOT EXISTS ")

	// add table name
	b.WriteString(s.tableName + " ")

	// add stable
	if s.sTableName != "" {
		b.WriteString("USING " + s.sTableName + " ")
	}

	// add tags
	if len(s.tags) > 0 {
		tagsKeys := make([]string, 0, len(s.tags))
		tagsValues := make([]any, 0, len(s.tags))
		for key, value := range s.tags {
			tagsKeys = append(tagsKeys, key)
			tagsValues = append(tagsValues, value)
		}
		b.WriteString("(" + strings.Join(tagsKeys, ", ") + ") ")

		b.WriteString("TAGS (")
		for index, value := range tagsValues {
			v, err := castType(value)
			if err != nil {
				return "", err
			}
			b.WriteString(v)

			if index != len(tagsValues)-1 {
				b.WriteString(", ")
			}
		}
		b.WriteString(")")
	}

	b.WriteString(";")

	return b.String(), nil
}

func (s CreateTableBuilder) validate() error {
	if s.tableName == "" {
		return ErrEmptyTableName
	}

	return nil
}
