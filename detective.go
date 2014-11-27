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

func fanin(chans []<-chan plugins.Result) chan plugins.Result {
	out := make(chan plugins.Result)
	for _, channel := range chans {
		go func(in <-chan plugins.Result) {
			for result := range in {
				out <- result
			}
		}(channel)
	}
	return out
}

func CollectData() map[string]interface{}{
	var(
		wg sync.WaitGroup
		channels []<-chan plugins.Result
		errchannels []<-chan error
	)
	data := make(map[string]interface{})
	done := make(chan bool)
	wg.Add(PluginReg.plugins.Len())
	for p := PluginReg.plugins.Front(); p != nil; p = p.Next() {
		plugin := p.Value.(plugins.Plugin)
		c, e := plugins.CollectorWrapper(done, &wg, plugin)
		channels = append(channels, c)
		errchannels = append(errchannels, e)
	}
	out := fanin(channels)
	return data
}
