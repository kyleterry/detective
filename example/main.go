package main

import "fmt"
import "github.com/kyleterry/detective"

func main() {
	detective.Init(true) // true enables debug logging
	d := detective.CollectData()
	platform := d["platform"].(map[string]interface{})
	fmt.Printf("%+v", d)
	version := platform["version"].(string)
	fmt.Printf("%s\n", version)
}
