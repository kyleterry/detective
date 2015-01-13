// +build darwin
package plugins

import (
	"log"
	"os/exec"
	"strconv"
)

type Cpu struct {
	Name string
}

func GetCpuMetric(metric string) (string, error) {

	cmd := exec.Command("sysctl", "-n", metric)
	result, err := cmd.Output()
	return string(result), err

}

// Cpu.CollectData returns cpu information about the system.
// Returns a string and a map[string]MetricValue.
func (m *Cpu) CollectData() (string, map[string]*MetricValue, error) {

	identifiers := map[string]string{
		"real":       "hw.physicalcpu",
		"total":      "hw.logicalcpu",
		"mhz":        "hw.cpufrequency",
		"vendor_id":  "machdep.cpu.brand_string",
		"model_name": "machdep.cpu.model",
		"model":      "machdep.cpu.family",
		"family":     "machdep.cpu.stepping",
		"stepping":   "machdep.cpu.features",
	}

	data := make(map[string]*MetricValue)

	for metric, sysMetric := range identifiers {

		val, err := GetCpuMetric(sysMetric)

		if err != nil {
			log.Fatal(err)
		}

		data[metric] = &MetricValue{val}

	}

	// Convert the mhz field from hertz to mhz
	mhz, err := strconv.Atoi((*data["mhz"]).Val)
	if err != nil {
		log.Fatal(err)
	}
	(*data["mhz"]).Val = string(mhz / 1000)

	return m.Name, data, nil
}
