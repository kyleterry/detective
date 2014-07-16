package detective

import (
	"fmt"
	"runtime"
	"github.com/kyleterry/go-detective/plugins"
)

func getOSType() string {
	return runtime.GOOS
}

func CollectData() map[string]interface{}{
	registerPlugins()
	data := make(map[string]interface{})
	switch getOSType() {
		case "linux":
			for lp := linuxPlugins.plugins.Front(); lp != nil; lp = lp.Next() {
				plug := lp.Value.(plugins.Plugin)
				name, d := plug.CollectData()
				data[name] = d
			}
	}
	fmt.Printf("%+v", data)
	return data
}
