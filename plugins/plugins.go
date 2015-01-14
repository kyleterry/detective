package plugins

type MetricValue struct {
	Val string
}

// Collection is a group of metrics namespaced by their plugin name.
type Collection struct {
	Items map[string]Result
}

// Result is a struct that holds a map of metrics for a specific plugin.
type Result struct {
	PluginName string
	Metrics    map[string]*MetricValue
}

// DataCollector's must define a CollectData() method.
//
// `
// type MemoryPlugin struct {}
//
// func (plug MemoryPlugin) CollectData() (string, map[string]*MetricValue, error) { ... }
// `
type DataCollector interface {
	CollectData() (string, map[string]*MetricValue, error)
}

// Returns a Collection instance
func NewCollection() Collection {
	return Collection{make(map[string]Result)}
}

// CollectorWrapper wraps plugins to create channels and call CollectData as a
// goroutine.
//
// Returns a result channel and an error channel
func CollectorWrapper(done <-chan bool, plug DataCollector) (<-chan Result, <-chan error) {
	c := make(chan Result)
	errc := make(chan error, 1)

	go func() {
		defer close(c)
		defer close(errc)

		name, metrics, err := plug.CollectData()
		if err != nil {
			errc <- err
		}

		select {
		case c <- Result{name, metrics}:
		case <-done:
		}
	}()

	return c, errc
}
