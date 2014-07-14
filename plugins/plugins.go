package plugins

type Plugin interface {
	OsType() string
	Run()
}
