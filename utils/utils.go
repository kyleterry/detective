package utils

import (
	"os"
	"os/exec"
	"io/ioutil"
	"github.com/kyleterry/go-detective"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("detective")

type cacheStore struct {
	lsbCache string
}

var cache cacheStore

func GetRawLSB() (string, error) {
	if detective.GetBirdsEyeOSType() != "linux" {
		log.Fatal("Can't get LSB information from non Linux system")
	}

	if cache.lsbCache == "" {
		if _, err := os.Stat("/etc/lsb-release"); err == nil {
			lsb_output, err := ioutil.ReadFile("/etc/lsb-release")
			if err != nil {
				return "", err
			}
			cache.lsbCache = string(lsb_output[:])
		} else {
			cmd := exec.Command("lsb_release", "-a")
			lsb_output, err := cmd.Output()
			if err != nil {
				return "", err
			}
			cache.lsbCache = string(lsb_output[:])
		}
	}
	return cache.lsbCache, nil
}
