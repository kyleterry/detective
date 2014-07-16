package utils

import (
	"os"
	"os/exec"
	"io/ioutil"
)

type cacheStore struct {
	lsbCache string
}

var cache cacheStore

func GetRawLSB() (string, error) {
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
