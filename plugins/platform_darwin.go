// +build darwin
package plugins

import (
	"os/exec"
)

type Platform struct {
	Name string
}

func (p *Platform) CollectData() (string, map[string]*MetricValue, error) {
	data := make(map[string]*MetricValue)
	cmd := exec.Command("sw_vers", "-productVersion")
	ver, err := cmd.Output()
	if err != nil {
		return "", nil, err
	}
	data["version"] = &MetricValue{string(ver)}
	return p.Name, data, nil
}
