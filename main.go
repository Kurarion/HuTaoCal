package main

import (
	"flag"
	"fmt"
	core "huTao/core"
	other "huTao/other"
)

var (
	pathJSON  string
	otherFlag bool
)

func init() {
	defaultJSONPath := "data.json"
	flag.StringVar(&pathJSON, "i", defaultJSONPath, "JSONファイルパス")
	flag.BoolVar(&otherFlag, "other", false, "other")
}

func main() {
	flag.Parse()
	// fmt.Printf("Damage: %v", core.CalDamage(pathJSON))
	// fmt.Println(core.GenerateJSON())
	if otherFlag {
		other.TestArtifact()
	} else {
		fmt.Println(core.GenerateJSON(core.CalDamage(pathJSON)))
	}
}
