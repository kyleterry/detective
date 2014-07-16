package plugins

type Plugin interface {
	OsType() string
	CollectData() (string, map[string]interface{})
}
