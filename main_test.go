package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	mapNode, err := ParseFileToMap("/Users/atom/scratches/walmart-rules.json")
	if err != nil {
		panic(err)
	}
	get := mapNode.GetValue("info", "title").ToString()
	fmt.Println(get)

	servers := mapNode.GetArray("servers")
	fmt.Println(servers)
	items := mapNode.GetArray("servers").Items()
	for _, item := range items {
		fmt.Println(fmt.Sprintf("%T", item))

	}
}
