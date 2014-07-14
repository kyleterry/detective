package linplug

type LinuxPlatform struct {
	Name string
}

func (self LinuxPlatform) OsType() string {
	return "linux"
}

func (self LinuxPlatform) CollectData() map[string]interface{} {
	data := make(map[string]interface{})
	return data
}
