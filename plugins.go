package detective

import (
	"container/list"
	"github.com/kyleterry/go-detective/plugins"
	"github.com/kyleterry/go-detective/plugins/linux"
)


type osPluginsRegistry struct {
	plugins *list.List
}

func (self *osPluginsRegistry) RegisterPlugin(p plugins.Plugin) {
	self.plugins.PushBack(p)
}

var linuxPlugins osPluginsRegistry
var osxPlugins osPluginsRegistry
var windowsPlugins osPluginsRegistry

func registerAllPlugins() {
	registerLinuxPlugins()
	registerOsxplugins()
	registerWindowsPlugins()
}

func registerLinuxPlugins() {
	linuxPlugins.plugins = list.New()
	platform := linplug.LinuxPlatform{"platform"}
	linuxPlugins.RegisterPlugin(platform)
}

func registerOsxplugins() {
	osxPlugins.plugins = list.New()
}

func registerWindowsPlugins() {
	windowsPlugins.plugins = list.New()
}
