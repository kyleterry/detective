package detective

import (
	"container/list"
	"github.com/kyleterry/detective/plugins"
	"github.com/kyleterry/detective/plugins/linux"
	"github.com/kyleterry/detective/plugins/osx"
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
	registerOSXPlugins()
	registerWindowsPlugins()
}

func registerLinuxPlugins() {
	linuxPlugins.plugins = list.New()
	platform := linux.LinuxPlatform{"platform"}
	linuxPlugins.RegisterPlugin(platform)
	memory := linux.LinuxMemory{"memory"}
	linuxPlugins.RegisterPlugin(memory)
}

func registerOSXPlugins() {
	osxPlugins.plugins = list.New()
	platform := osx.OSXPlatform{"platform"}
	osxPlugins.RegisterPlugin(platform)
}

func registerWindowsPlugins() {
	windowsPlugins.plugins = list.New()
}
