package goson

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	jsonNode, err := ParseFile("test-data/data-1.json")
	if err != nil {
		panic(err)
	}
	{
		popup := AsMap(jsonNode).GetMap("menu", "popup")
		assert.NotNil(t, popup)
		assert.IsType(t, &MapNode{}, popup)
	}
	{
		popup := AsMap(jsonNode).GetMap("menu", "popup", "items")
		assert.Nil(t, popup)
	}
	{
		items := AsMap(jsonNode).GetArray("menu", "popup", "items")
		assert.NotNil(t, items)
		assert.IsType(t, &ArrayNode{}, items)
		assert.Equal(t, 3, len(items.Items()))
		assert.Equal(t, "Open", AsValue(items.Items()[1]).String())
	}

}

func TestValue(t *testing.T) {
	mapNode, err := ParseFileToMap("test-data/data-1.json")
	if err != nil {
		panic(err)
	}
	{
		get := mapNode.GetValue("menu", "id")
		assert.NotNil(t, get)
		assert.IsType(t, &ValueNode{}, get)
		assert.Equal(t, "file", get.String())
		assert.Equal(t, "file", get.ToString())
	}
	{
		get := mapNode.GetValue("menu", "order")
		assert.NotNil(t, get)
		assert.IsType(t, &ValueNode{}, get)
		assert.Equal(t, "", get.String())
		assert.Equal(t, "10", get.ToString())
	}
	{
		get := mapNode.GetString("menu", "id")
		assert.NotNil(t, get)
		assert.IsType(t, "", get)
		assert.Equal(t, "file", get)
	}
	{
		get := mapNode.GetToString("menu", "id")
		assert.NotNil(t, get)
		assert.IsType(t, "", get)
		assert.Equal(t, "file", get)
	}
	{
		get := mapNode.GetString("menu", "order")
		assert.NotNil(t, get)
		assert.IsType(t, "", get)
		assert.Equal(t, "", get)
	}
	{
		get := mapNode.GetToString("menu", "order")
		assert.NotNil(t, get)
		assert.IsType(t, "", get)
		assert.Equal(t, "10", get)
	}
}
