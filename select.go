package tsbuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ TdEngineSQLBuilder = (*SelectBuilder)(nil)

type SelectBuilder struct {
	columns         []string
	from            string
	whereConditions []string
	groupBy         string
	partitionBy     string
	orderBy         string
	limit           *uint32
	slimit          *uint32
	offset          *uint32
	soffset         *uint32
}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{
		columns:         make([]string, 0),
		whereConditions: make([]string, 0),
	}
}

func (s *SelectBuilder) Columns(columns ...string) *SelectBuilder {
	s.columns = columns
	return s
}

func (s *SelectBuilder) From(from string) *SelectBuilder {
	s.from = from
	return s
}

func (s *SelectBuilder) Where(conditions ...string) *SelectBuilder {
	s.whereConditions = append(s.whereConditions, conditions...)
	return s
}

func (s *SelectBuilder) GroupBy(value string) *SelectBuilder {
	s.groupBy = value
	return s
}

func (s *SelectBuilder) PartitionBy(value string) *SelectBuilder {
	s.partitionBy = value
	return s
}

func (s *SelectBuilder) OrderBy(value string) *SelectBuilder {
	s.orderBy = value
	return s
}

func (s *SelectBuilder) Limit(value *uint32) *SelectBuilder {
	s.limit = value
	return s
}

func (s *SelectBuilder) SLimit(value *uint32) *SelectBuilder {
	s.slimit = value
	return s
}

func (s *SelectBuilder) Offset(value *uint32) *SelectBuilder {
	s.offset = value
	return s
}

func (s *SelectBuilder) SOffset(value *uint32) *SelectBuilder {
	s.soffset = value
	return s
}

func (s *SelectBuilder) Build() (string, error) {
	if err := s.validate(); err != nil {
		return "", fmt.Errorf("validate error: %w", err)
	}

	b := bytes.NewBuffer([]byte{})
	b.WriteString("SELECT ")

	// add columns
	b.WriteString(strings.Join(s.columns, ", "))
	b.WriteString(" ")

	// add from
	b.WriteString("FROM " + s.from + " ")

	// add where conditions
	if len(s.whereConditions) > 0 {
		b.WriteString("WHERE ")
		b.WriteString(strings.Join(s.whereConditions, " AND "))
	}

	// add group by
	if s.groupBy != "" {
		b.WriteString(" GROUP BY " + s.groupBy + " ")
	}

	// add partition by
	if s.partitionBy != "" {
		b.WriteString(" PARTITION BY " + s.partitionBy + " ")
	}

	// add order by
	if s.orderBy != "" {
		b.WriteString(" ORDER BY " + s.orderBy + " ")
	}

	// add limit
	if s.limit != nil {
		b.WriteString(fmt.Sprintf("LIMIT %d", *s.limit) + " ")
	}

	// add slimit
	if s.slimit != nil {
		b.WriteString(fmt.Sprintf("SLIMIT %d", *s.slimit) + " ")
	}

	// add offset
	if s.offset != nil {
		fmt.Fprintf(b, "OFFSET %d", *s.offset)
	}

	// add soffset
	if s.soffset != nil {
		if s.offset != nil {
			b.WriteString(" ")
		}
		fmt.Fprintf(b, "SOFFSET %d", *s.soffset)
	}

	b.WriteString(";")

	return b.String(), nil
}

func (s *SelectBuilder) validate() error {
	if s.from == "" {
		return ErrEmptyFrom
	}

	if len(s.columns) == 0 {
		return ErrEmptyColumns
	}

	return nil
}
