package osx

import (
	"os/exec"
	"github.com/kyleterry/detective/plugins"
)

type OSXPlatform struct {
	Name string
}

func (self OSXPlatform) OsType() string {
	return "darwin"
}

func (self OSXPlatform) CollectData() (string, map[string]*plugins.MetricValue) {
	data := make(map[string]*plugins.MetricValue)
	cmd := exec.Command("sw_vers", "-productVersion")
	ver, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	data["version"] = &plugins.MetricValue{string(ver)}
	return self.Name, data
}
