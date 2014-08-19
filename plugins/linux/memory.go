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
		switch section {
		case "MemTotal":
			value_int, err = strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			value_int = value_int / 1024 / 1024
			data[section] = value_int
		}
	}
	return self.Name, data
}

func SplitMeminfoLine(line string) (string, string) {
	var idx int
	idx = strings.Index(line, " ")
	if idx == -1 {
		log.Fatal("Can't parse meminfo file")
	}
	section := string(line[:idx - 1])
	line = line[idx:]
	idx = strings.Index(line, " ")
	value := string(line[:idx])
	return section, value
}
