package plugins

type MetricValue struct {
	Val string
}

type Collection struct {
	Items map[string]*MetricValue
}

type Result struct {
	Name    string
	Metrics map[string]*MetricValue
}

type DataCollector interface {
	CollectData() (string, map[string]*MetricValue, error)
}

type Plugin interface {
	DataCollector
}

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
