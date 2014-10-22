// +build darwin
package plugins

type Memory struct {
	Name string
}

const (
	MeminfoFile = "/proc/meminfo"
)

func (m *Memory) OsType() string {
	return "linux"
}

// Memory.CollectData returns memory information about the system.
// Values are in KB.
// Returns a string and a map[string]MetricValue.
func (m *Memory) CollectData() (string, map[string]*MetricValue) {
	data := make(map[string]*MetricValue)
	return m.Name, data
}
