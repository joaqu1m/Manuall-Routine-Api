package types

type Field struct {
	DbCol  string
	Opts   map[string]interface{}
	Format func(string) string
}
