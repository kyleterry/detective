package main

import "fmt"
import "github.com/kyleterry/detective"

func main() {
	detective.Init(true) // true enables debug logging
	d := detective.CollectData()
	fmt.Printf("%+v", d)
}
