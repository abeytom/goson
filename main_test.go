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

func TestFind(t *testing.T) {
	n, err := ParseFile("test-data/data-1.json")
	assert.Nil(t, err)
	{
		node := Find(n, "menu")
		assert.NotNil(t, AsMap(node))
		assert.Equal(t, "file", AsMap(node).GetString("id"))
	}
	{
		node := Find(n, "popup")
		assert.NotNil(t, AsMap(node))
	}
	{
		node := Find(n, "menuitem")
		assert.NotNil(t, AsArray(node))
	}
	{
		node := Find(n, "menuitem", "value")
		assert.Equal(t, "New", AsValue(node).String())
	}
	{
		nodes := FindAll(n, "menuitem", "value")
		assert.Equal(t, 5, len(nodes))
		assert.Equal(t, "New", AsValue(nodes[0]).String())
		assert.Equal(t, "Open", AsValue(nodes[1]).String())
		assert.Equal(t, "Save", AsValue(nodes[2]).String())
		//fixme we have an map ordering problem here. revisit later
		assert.Equal(t, "SaveDocChild", AsValue(nodes[3]).String())
		assert.Equal(t, "SaveDocGChild", AsValue(nodes[4]).String())

	}
	{
		nodes := FindAll(n, "menuitem", "onclick")
		assert.Equal(t, 4, len(nodes))
	}
}
