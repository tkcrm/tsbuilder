package tdenginebuilder

type tdenginebuilder interface {
	Build() (sql string, err error)
}
