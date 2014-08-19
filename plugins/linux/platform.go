package linux

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/kyleterry/detective/utils"
)


type LinuxPlatform struct {
	Name string
}

func (self LinuxPlatform) OsType() string {
	return "linux"
}

func (self LinuxPlatform) CollectData() (string, map[string]interface{}) {
	data := make(map[string]interface{})
	rawLsb, err := utils.GetRawLSB()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Detecting distro")
	if _, err := os.Stat("/etc/debian_version"); err == nil {
		log.Debug("Found Debian")
		deb_re := regexp.MustCompile("/Ubuntu/")
		if deb_re.MatchString(rawLsb) {
			data["version"] = "ubuntu"
		} else {
			data["distro"] = "debian"
		}
		deb_ver, _ := ioutil.ReadFile("/etc/debian_version")
		data["version"] = string(deb_ver)
	} else if _, err := os.Stat("/etc/gentoo-release"); err == nil {
		data["distro"] = "gentoo"
		gentoo_ver, _ := ioutil.ReadFile("/etc/gentoo-release")
		data["version"] = string(gentoo_ver)
	} else if _, err := os.Stat("/etc/arch-release"); err == nil {
		data["distro"] = "arch"
		data["version"] = ""
	} else if _, err := os.Stat("/etc/slackware-version"); err == nil {
		data["distro"] = "slackware"
		slack_ver, _ := ioutil.ReadFile("/etc/slackware-version")
		data["version"] = string(slack_ver)
	}

	return self.Name, data
}
