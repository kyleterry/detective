package detective

import (
	"fmt"
	stdlog "log"
	"os"
	"runtime"
	"github.com/kyleterry/go-detective/plugins"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("detective")

func Init() {
	logBackend := logging.NewLogBackend(os.Stdout, "", stdlog.LstdFlags)
	logBackend.Color = true
	logging.SetBackend(logBackend)
}

func GetBirdsEyeOSType() string {
	return runtime.GOOS
}

func CollectData() map[string]interface{}{
	data := make(map[string]interface{})
	switch GetBirdsEyeOSType() {
		case "linux":
			registerLinuxPlugins()
			for lp := linuxPlugins.plugins.Front(); lp != nil; lp = lp.Next() {
				plug := lp.Value.(plugins.Plugin)
				name, d := plug.CollectData()
				data[name] = d
			}
	}
	fmt.Printf("%+v", data)
	return data
}
