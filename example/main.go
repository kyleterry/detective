package main

import "fmt"
import "github.com/kyleterry/detective"

func main() {
	detective.Init() // true enables debug logging
	results := detective.CollectAllMetrics()
	for plugin, result := range results.Items {
		fmt.Printf("%#v\n", plugin)
		fmt.Printf("%#v\n", result)
	}
}
