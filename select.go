package tdenginebuilder

import (
	"bytes"
	"fmt"
	"strings"
)

var _ tdenginebuilder = (*selectBuilder)(nil)

type selectBuilder struct {
	columns         []string
	from            string
	whereConditions []string
	groupBy         string
	partitionBy     string
	orderBy         string
	limit           *uint64
	slimit          *uint64
	offset          *uint64
	soffset         *uint64
}

func NewSelectBuilder() *selectBuilder {
	return &selectBuilder{
		columns:         make([]string, 0),
		whereConditions: make([]string, 0),
	}
}

func (s *selectBuilder) Columns(columns ...string) *selectBuilder {
	s.columns = columns
	return s
}

func (s *selectBuilder) From(from string) *selectBuilder {
	s.from = from
	return s
}

func (s *selectBuilder) Where(conditions ...string) *selectBuilder {
	s.whereConditions = append(s.whereConditions, conditions...)
	return s
}

func (s *selectBuilder) GroupBy(value string) *selectBuilder {
	s.groupBy = value
	return s
}

func (s *selectBuilder) PartitionBy(value string) *selectBuilder {
	s.partitionBy = value
	return s
}

func (s *selectBuilder) OrderBy(value string) *selectBuilder {
	s.orderBy = value
	return s
}

func (s *selectBuilder) Limit(value *uint64) *selectBuilder {
	s.limit = value
	return s
}

func (s *selectBuilder) SLimit(value *uint64) *selectBuilder {
	s.slimit = value
	return s
}

func (s *selectBuilder) Offset(value *uint64) *selectBuilder {
	s.offset = value
	return s
}

func (s *selectBuilder) SOffset(value *uint64) *selectBuilder {
	s.soffset = value
	return s
}

func (s *selectBuilder) Build() (string, error) {
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
		b.WriteString(s.groupBy + " ")
	}

	// add partition by
	if s.partitionBy != "" {
		b.WriteString(s.partitionBy + " ")
	}

	// add order by
	if s.orderBy != "" {
		b.WriteString(s.orderBy + " ")
	}

	// add limit
	if s.limit != nil {
		b.WriteString(fmt.Sprintf("%d", *s.limit) + " ")
	}

	// add slimit
	if s.slimit != nil {
		b.WriteString(fmt.Sprintf("%d", *s.slimit) + " ")
	}

	// add offset
	if s.offset != nil {
		b.WriteString(fmt.Sprintf("%d", *s.offset) + " ")
	}

	// add soffset
	if s.soffset != nil {
		b.WriteString(fmt.Sprintf("%d", *s.soffset) + " ")
	}

	b.WriteString(";")

	return b.String(), nil
}

func (s *selectBuilder) validate() error {
	if s.from == "" {
		return ErrEmptyFrom
	}

	if len(s.columns) == 0 {
		return ErrEmptyColumns
	}

	return nil
}
