package tsbuilder

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var _ tdEngineSqlBuilder = (*insertBuilder)(nil)

type insertTableData struct {
	tableName string
	using     string
	columns   []string
	values    [][]any
	tags      map[string]any
}

type insertBuilder struct {
	tablesData []*insertTableData
}

func NewInsertBuilder() *insertBuilder {
	return &insertBuilder{
		tablesData: make([]*insertTableData, 0),
	}
}

func (s *insertBuilder) AddTable(tableName string) *insertTableData {
	for idx, table := range s.tablesData {
		if table.tableName == tableName {
			s.tablesData = append(s.tablesData[:idx], s.tablesData[idx+1:]...)
		}
	}

	tData := &insertTableData{
		tableName: tableName,
		tags:      make(map[string]any),
		columns:   make([]string, 0),
		values:    make([][]any, 0),
	}

	s.tablesData = append(s.tablesData, tData)

	return tData
}

func (s *insertTableData) Using(v string) *insertTableData {
	s.using = v
	return s
}

func (s *insertTableData) Columns(columns ...string) *insertTableData {
	s.columns = columns
	return s
}

func (s *insertTableData) Values(values ...any) *insertTableData {
	s.values = append(s.values, values)
	return s
}

func (s *insertTableData) Tags(tags map[string]any) *insertTableData {
	if tags == nil {
		return s
	}

	s.tags = tags
	return s
}

func (s *insertBuilder) Build() (string, error) {
	b := bytes.NewBuffer([]byte{})
	b.WriteString("INSERT INTO ")

	for tIndex, table := range s.tablesData {
		if err := table.validate(); err != nil {
			return "", fmt.Errorf("validate error: %w", err)
		}

		// add table name
		b.WriteString(table.tableName + " ")

		// add using
		if table.using != "" {
			b.WriteString("USING " + table.using + " ")
		}

		// add tags
		if len(table.tags) > 0 {
			tagsKeys := make([]string, 0, len(table.tags))
			tagsValues := make([]any, 0, len(table.tags))
			for key, value := range table.tags {
				tagsKeys = append(tagsKeys, key)
				tagsValues = append(tagsValues, value)
			}

			b.WriteString("(")
			for index, value := range tagsKeys {
				b.WriteString(value)

				if index != len(tagsKeys)-1 {
					b.WriteString(", ")
				}
			}
			b.WriteString(") ")

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
			b.WriteString(") ")
		}

		// add columns
		if len(table.columns) > 0 {
			b.WriteString("(" + strings.Join(table.columns, ", ") + ") ")
		}

		// add values
		b.WriteString("VALUES ")
		for index, valuesGroup := range table.values {
			if len(table.columns) != len(valuesGroup) {
				return "", errors.New("columns should be equal values")
			}

			b.WriteString("(")
			for vIndex, value := range valuesGroup {
				v, err := castType(value)
				if err != nil {
					return "", err
				}
				b.WriteString(v)
				if vIndex != len(valuesGroup)-1 {
					b.WriteString(", ")
				}
			}
			b.WriteString(")")
			if index != len(table.values)-1 {
				b.WriteString(" ")
			}
		}

		if tIndex != len(s.tablesData)-1 {
			b.WriteString(", ")
		}
	}

	b.WriteString(";")

	return b.String(), nil
}

func (s *insertTableData) validate() error {
	if s.tableName == "" {
		return ErrEmptyTableName
	}

	if len(s.columns) == 0 {
		return ErrEmptyColumns
	}

	if len(s.values) == 0 {
		return ErrEmptyValues
	}

	return nil
}
