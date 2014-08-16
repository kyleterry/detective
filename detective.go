package detective

import (
	stdlog "log"
	"os"
	"github.com/kyleterry/detective/plugins"
	"github.com/kyleterry/detective/utils"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("detective")

func Init() {
	logBackend := logging.NewLogBackend(os.Stdout, "", stdlog.LstdFlags)
	logBackend.Color = true
	logging.SetBackend(logBackend)
}

func CollectData() map[string]interface{}{
	data := make(map[string]interface{})
	switch utils.GetBirdsEyeOSType() {
		case "linux":
			registerLinuxPlugins()
			for lp := linuxPlugins.plugins.Front(); lp != nil; lp = lp.Next() {
				plug := lp.Value.(plugins.Plugin)
				name, d := plug.CollectData()
				data[name] = d
			}
		case "darwin":
			registerOSXPlugins()
			for op := osxPlugins.plugins.Front(); op != nil; op = op.Next() {
				plug := op.Value.(plugins.Plugin)
				name, d := plug.CollectData()
				data[name] = d
			}
			
	}
	return data
}
