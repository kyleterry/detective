package detective

import (
	stdlog "log"
	"os"
	"github.com/kyleterry/detective/plugins"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("detective")

func Init(debug bool) {
	logBackend := logging.NewLogBackend(os.Stdout, "", stdlog.LstdFlags)
	logBackend.Color = true
	logging.SetBackend(logBackend)
}

func CollectData() map[string]interface{}{
	data := make(map[string]interface{})
	for p := PluginReg.plugins.Front(); p != nil; p = p.Next() {
		plugin := p.Value.(plugins.Plugin)
		name, d := plugin.CollectData()
		if d == nil {
			log.Error("Error collecting: %s", name)
		}
		data[name] = d
	}
	return data
}
