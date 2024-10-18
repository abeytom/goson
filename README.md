# goson

Introduction
------------

This is just a convenience wrapper around the go map hierarchy. Makes it easier process/iterate/update `json` or `yaml`
documents.

Usage
-----

```go
package main

import (
	"encoding/json"
	"github.com/abeytom/goson"
)

func main() {
	// Read json file
	mapNode, err := goson.ParseFileToMap("test-data/data-1.json")
	// or Read yaml file
	mapNode, err := goson.ParseYamlFileToMap("test-data/data-1.yml")

	// Get the nested values
	strValue := mapNode.GetString("menu", "id")                         // gets "file"
	strValue := mapNode.GetToString("menu", "order")                    // gets "10"
	floatValue := mapNode.GetValue("menu", "order").Value().(float64)   // get 10
	emptyStr := mapNode.GetString("menu", "not-exists", "not-exists-2") // gets ""

	// iterate over an array
	array := mapNode.GetArray("menu", "popup", "menuitem")
	for _, node := range array.ItemsAsMap() {
		strValue := node.GetString("value")
	}

	// iterate over map keys
	for k, node := range mapNode.GetMap("menu").EntriesAsMap() {
		strValue := node.GetToString("order") // gets "10"
	}

	// update
	var val interface{} // = ...
	mapNode.GetMap("menu").Set("key", val)

	// use json.Marshall or yaml.Marshall to split the contents
	b, err := json.Marshal(mapNode.Object)
	// or 
	b, err := json.Marshal(mapNode.GetMap("menu").Object)
}
```

See the [main_test.go](main_test.go) for more usage



