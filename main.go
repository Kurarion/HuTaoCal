package main

import (
	"flag"
	core "huTao/core"
)

var (
	pathJSON string
)

func init() {
	defaultJSONPath := "data.json"
	flag.StringVar(&pathJSON, "i", defaultJSONPath, "JSONファイルパス")
}

func main() {
	flag.Parse()
	// fmt.Printf("Damage: %v", core.CalDamage(pathJSON))
	// fmt.Println(core.GenerateJSON())
	core.CalDamage(pathJSON)
}
