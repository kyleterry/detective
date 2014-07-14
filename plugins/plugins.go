package plugins

type Plugin interface {
	OsType() string
	CollectData() map[string]interface{}
}
