package tsbuilder

type tdEngineSqlBuilder interface {
	Build() (sql string, err error)
}
