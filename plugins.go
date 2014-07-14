package detective

import (
	"container/list"
	"github.com/kyleterry/go-detective/plugins"
)


type osPluginsRegistry struct {
	plugins *list.List
}

func (self *osPluginsRegistry) AddPlugin(p plugins.Plugin) {
	self.plugins.PushBack(p)
}

var linuxPlugins osPluginsRegistry
var osxPlugins osPluginsRegistry
var windowsPlugins osPluginsRegistry

func registerPlugins() {
	// Try to make this more lazy. Maybe call after detecting the OS?
	registerLinuxPlugins()
	registerOsxplugins()
	registerWindowsPlugins()
}

func registerLinuxPlugins() {
	linuxPlugins.plugins = list.New()
	platform := make(plugins.LinuxPlatform)
	linuxPlugins.AddPlugin(plugins.LinuxPlatform)
}

func registerOsxplugins() {
	osxPlugins.plugins = list.New()
}

func registerWindowsPlugins() {
	windowsPlugins.plugins = list.New()
}
