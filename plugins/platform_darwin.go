// +build darwin
package plugins

import (
	"log"
	"os/exec"
)

type Platform struct {
	Name string
}

func (p *Platform) OsType() string {
	return "darwin"
}

func (p *Platform) CollectData() (string, map[string]*MetricValue) {
	data := make(map[string]*MetricValue)
	cmd := exec.Command("sw_vers", "-productVersion")
	ver, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	data["version"] = &MetricValue{string(ver)}
	return p.Name, data
}
