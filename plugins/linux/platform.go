package linux

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/kyleterry/detective/utils"
	"github.com/kyleterry/detective/plugins"
)


type LinuxPlatform struct {
	Name string
}

func (self LinuxPlatform) OsType() string {
	return "linux"
}

func (self LinuxPlatform) CollectData() (string, map[string]*plugins.MetricValue) {
	data := make(map[string]*plugins.MetricValue)
	var version []byte
	rawLsb, err := utils.GetRawLSB()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Detecting distro")
	if _, err := os.Stat("/etc/debian_version"); err == nil {
		log.Debug("Found Debian")
		deb_re := regexp.MustCompile("/Ubuntu/")
		if deb_re.MatchString(rawLsb) {
			data["version"] = &plugins.MetricValue{string("ubuntu")}
		} else {
			data["distro"] = &plugins.MetricValue{string("debian")}
		}
		version, _ = ioutil.ReadFile("/etc/debian_version")
	} else if _, err := os.Stat("/etc/gentoo-release"); err == nil {
		data["distro"] = &plugins.MetricValue{string("gentoo")}
		version, _ = ioutil.ReadFile("/etc/gentoo-release")
	} else if _, err := os.Stat("/etc/arch-release"); err == nil {
		data["distro"] = &plugins.MetricValue{string("arch")}
		version = nil
	} else if _, err := os.Stat("/etc/slackware-version"); err == nil {
		data["distro"] = &plugins.MetricValue{string("slackware")}
		version, _ = ioutil.ReadFile("/etc/slackware-version")
	}
	data["version"] = &plugins.MetricValue{string(strings.TrimSpace(string(version)))}

	return self.Name, data
}
