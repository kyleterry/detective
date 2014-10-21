package detective

import (
	"container/list"
	"github.com/kyleterry/detective/plugins"
)

type PluginsRegistry struct {
	plugins *list.List
}

var PluginReg PluginsRegistry

func (self *PluginsRegistry) RegisterPlugin(p plugins.Plugin) {
	self.plugins.PushBack(p)
}

func init() {

	PluginReg.plugins = list.New()
	platform := plugins.Platform{"platform"}
	PluginReg.RegisterPlugin(platform)
	memory := plugins.Memory{"memory"}
	PluginReg.RegisterPlugin(memory)
}
