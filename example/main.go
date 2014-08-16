package main

import "fmt"
import "github.com/kyleterry/detective"

func main() {
	detective.Init()
	d := detective.CollectData()
	platform := d["platform"].(map[string]interface{})
	version := platform["version"].(string)
	fmt.Printf("%s\n", version)
}
