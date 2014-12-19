package detective

import (
	stdlog "log"
	"os"
	"sync"
	"github.com/kyleterry/detective/plugins"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("detective")

func Init() {
	logBackend := logging.NewLogBackend(os.Stdout, "", stdlog.LstdFlags)
	logBackend.Color = true
	logging.SetBackend(logBackend)
}

// fanin will merge all the channels dedicated to collecting metrics into one
// channel that CollectAllMetrics() can recieve from.
// It returns a channel that all collectors will be redirected to.
func fanin(wg *sync.WaitGroup, chans []<-chan plugins.Result) chan plugins.Result {
	out := make(chan plugins.Result)
	for _, channel := range chans {
		go func(in <-chan plugins.Result) {
			for result := range in {
				out <- result
			}
			wg.Done()
		}(channel)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func CollectAllMetrics() map[string]plugins.Result{
	var(
		wg sync.WaitGroup
		channels []<-chan plugins.Result
		errchannels []<-chan error
	)
	data := make(map[string]plugins.Result)
	done := make(chan bool)
	wg.Add(PluginReg.plugins.Len())
	for p := PluginReg.plugins.Front(); p != nil; p = p.Next() {
		plugin := p.Value.(plugins.Plugin)
		c, e := plugins.CollectorWrapper(done, plugin)
		channels = append(channels, c)
		errchannels = append(errchannels, e)
	}
	out := fanin(&wg, channels)
	for result := range out {
		data[result.Name] = result
	}
	return data
}
