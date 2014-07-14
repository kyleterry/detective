package detective

import (
	"runtime"
	"github.com/kyleterry/go-detective/plugins"
)

func getOSType() string {
	return runtime.GOOS
}

func CollectData() map[string]interface{}{
	data := make(map[string]interface{})
	switch getOSType() {
		case "linux":
			for lp := linuxPlugins.plugins.Front(); lp != nil; lp = lp.Next() {
				plug := lp.Value.(plugins.Plugin)
				d := plug.CollectData()
				data[plug.Name] = d
			}
	}
	return data
}
