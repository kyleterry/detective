// +build darwin
package plugins

import (
	"log"
	"os/exec"
)

type Platform struct {
	Name string
}

func (self Platform) OsType() string {
	return "darwin"
}

func (self Platform) CollectData() (string, map[string]*MetricValue) {
	data := make(map[string]*MetricValue)
	cmd := exec.Command("sw_vers", "-productVersion")
	ver, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	data["version"] = &MetricValue{string(ver)}
	return self.Name, data
}
