package osxplug

import (
	"os/exec"
)

type OSXPlatform struct {
	Name string
}

func (self OSXPlatform) OsType() string {
	return "darwin"
}

func (self OSXPlatform) CollectData() (string, map[string]interface{}) {
	data := make(map[string]interface{})
	cmd := exec.Command("sw_vers", "-productVersion")
	ver, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	data["version"] = string(ver)
	return self.Name, data
}
