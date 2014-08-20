package linux

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

type LinuxMemory struct {
	Name string
}

const (
	MeminfoFile = "/proc/meminfo"
)

func (self LinuxMemory) OsType() string {
	return "linux"
}

// LinuxMemory.CollectData returns memory information about the system.
// Values are in KB.
// Returns a string and a map[string]interface{}.
func (self LinuxMemory) CollectData() (string, map[string]interface{}) {
	data := make(map[string]interface{})
	file, err := os.Open(MeminfoFile)
	if err != nil {
		log.Fatal(err)
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
		data[section] = value_int
	}
	return self.Name, data
}

// SplitMeminfoLine splits each line of /proc/meminfo into key/value pairs.
// Returns a string and a string pair.
func SplitMeminfoLine(line string) (string, string) {
	fields := strings.Fields(line)
	return string(fields[0][:len(fields[0])-1]), fields[1]
}
