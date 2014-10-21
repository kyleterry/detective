// +build darwin
package plugins

type Memory struct {
	Name string
}

const (
	MeminfoFile = "/proc/meminfo"
)

func (self Memory) OsType() string {
	return "linux"
}

// Memory.CollectData returns memory information about the system.
// Values are in KB.
// Returns a string and a map[string]MetricValue.
func (self Memory) CollectData() (string, map[string]*MetricValue) {
	data := make(map[string]*MetricValue)
	return self.Name, data
}
