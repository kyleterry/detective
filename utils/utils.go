package utils

import (
	"github.com/op/go-logging"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var log = logging.MustGetLogger("detective")

type cacheStore struct {
	lsbCache string
}

var cache cacheStore

func GetRawLSB() (string, error) {
	if GetBirdsEyeOSType() != "linux" {
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

func GetBirdsEyeOSType() string {
	return runtime.GOOS
}
