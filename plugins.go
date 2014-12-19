package detective

import (
	"container/list"
	"github.com/kyleterry/detective/plugins"
)

type PluginsRegistry struct {
	plugins *list.List
}

var RegisteredPlugins PluginsRegistry

// RegisterPlugin is used to add a plugin to the registry for collecting.
// The plugin must implement the `plugins.DataCollector` interface.
func (self *PluginsRegistry) RegisterPlugin(p plugins.DataCollector) {
	self.plugins.PushBack(p)
}

func init() {
	RegisteredPlugins.plugins = list.New()
	platform := &plugins.Platform{"platform"}
	RegisteredPlugins.RegisterPlugin(platform)
	memory := &plugins.Memory{"memory"}
	RegisteredPlugins.RegisterPlugin(memory)
}
