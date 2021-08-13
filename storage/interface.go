package storage

type Engine interface {
	Query(method string, path string) (ok bool, rule string, vars map[string]string)
	Store(method, path string)
}
