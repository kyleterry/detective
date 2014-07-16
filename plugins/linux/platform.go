package linplug

import (
	"os"
)

type LinuxPlatform struct {
	Name string
}

func (self LinuxPlatform) OsType() string {
	return "linux"
}

func (self LinuxPlatform) CollectData() (string, map[string]interface{}) {
	data := make(map[string]interface{})
	if _, err := os.Stat("/etc/debian-version"); err == nil{
		data["platform"] = "debian"
	} else if _, err := os.Stat("/etc/gentoo-release"); err == nil{
		data["platform"] = "gentoo"
	}

	return self.Name, data
}
