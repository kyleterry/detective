// +build linux
package plugins

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/kyleterry/detective/utils"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("detective")

type Platform struct {
	Name string
}

func (p *Platform) OsType() string {
	return "linux"
}

func (p *Platform) CollectData() (string, map[string]*MetricValue) {
	data := make(map[string]*MetricValue)
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
			data["version"] = &MetricValue{string("ubuntu")}
		} else {
			data["distro"] = &MetricValue{string("debian")}
		}
		version, _ = ioutil.ReadFile("/etc/debian_version")
	} else if _, err := os.Stat("/etc/gentoo-release"); err == nil {
		data["distro"] = &MetricValue{string("gentoo")}
		version, _ = ioutil.ReadFile("/etc/gentoo-release")
	} else if _, err := os.Stat("/etc/arch-release"); err == nil {
		data["distro"] = &MetricValue{string("arch")}
		version = nil
	} else if _, err := os.Stat("/etc/slackware-version"); err == nil {
		data["distro"] = &MetricValue{string("slackware")}
		version, _ = ioutil.ReadFile("/etc/slackware-version")
	}
	data["version"] = &MetricValue{string(strings.TrimSpace(string(version)))}

	return p.Name, data
}
