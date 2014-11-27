package plugins

import(
	"sync"
)

type MetricValue struct {
	Val string
}

type Collection struct {
	Items map[string]*MetricValue
}

type Result struct {
	name    string
	results map[string]*MetricValue
}

type CollectDataer interface {
	CollectData() (string, map[string]*MetricValue, error)
}

type Plugin interface {
	CollectDataer
}

func CollectorWrapper(done <-chan bool, wg *sync.WaitGroup, plug CollectDataer) (<-chan Result, <-chan error) {
	c := make(chan Result)
	errc := make(chan error, 1)

	go func() {
		defer close(c)
		defer close(errc)
		defer wg.Done()

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
