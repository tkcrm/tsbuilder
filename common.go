package tsbuilder

type TdEngineSQLBuilder interface {
	Build() (sql string, err error)
}
