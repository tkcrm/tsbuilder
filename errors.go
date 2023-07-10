package tsbuilder

import "errors"

var (
	ErrEmptyDatabaseName = errors.New("empty database name")
	ErrEmptySTableName   = errors.New("empty stable name")
	ErrEmptyTableName    = errors.New("empty table name")
	ErrEmptyTags         = errors.New("empty tags")
	ErrEmptyColumns      = errors.New("empty columns")
	ErrEmptyValues       = errors.New("empty values")
	ErrEmptyFrom         = errors.New("empty from clause")
)
