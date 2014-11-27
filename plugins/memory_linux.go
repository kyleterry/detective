// +build linux
package plugins

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Memory struct {
	Name string
}

const (
	MeminfoFile = "/proc/meminfo"
)

// Memory.CollectData returns memory information about the system.
// Values are in KB.
// Returns a string and a map[string]interface{}.
func (m *Memory) CollectData() (string, map[string]*MetricValue, error) {
	data := make(map[string]*MetricValue)
	file, err := os.Open(MeminfoFile)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var (
		section, value string
		value_int int
	)
	for scanner.Scan() {
		line := scanner.Text()
		section, value = SplitMeminfoLine(line)
		value_int, err = strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		data[section] = &MetricValue{string(value_int)}
	}
	return m.Name, data, nil
}

// SplitMeminfoLine splits each line of /proc/meminfo into key/value pairs.
// Returns a string and a string pair.
func SplitMeminfoLine(line string) (string, string) {
	fields := strings.Fields(line)
	return string(fields[0][:len(fields[0])-1]), fields[1]
}
