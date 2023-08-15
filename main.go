package goson

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

type JsonNode interface {
}

type MapNode struct {
	Object map[string]interface{}
}

type ArrayNode struct {
	Objects []interface{}
}

type ValueNode struct {
	Val interface{}
}

func (o *MapNode) GetMap(keys ...string) *MapNode {
	return asMapNode(o.Get(keys...))
}

func (o *MapNode) GetValue(keys ...string) *ValueNode {
	return asValueNode(o.Get(keys...))
}

func (o *MapNode) GetString(keys ...string) string {
	node := asValueNode(o.Get(keys...))
	if node == nil {
		return ""
	}
	return node.String()
}

func (o *MapNode) GetToString(keys ...string) string {
	node := asValueNode(o.Get(keys...))
	if node == nil {
		return ""
	}
	return node.ToString()
}

func (o *MapNode) GetArray(keys ...string) *ArrayNode {
	return asArrayNode(o.Get(keys...))
}

func (v *ValueNode) Value() interface{} {
	return v.Val
}

func (v *ValueNode) String() string {
	s, ok := v.Val.(string)
	if ok {
		return s
	} else {
		return ""
	}
}

func (v *ValueNode) ToString() string {
	return fmt.Sprintf("%v", v.Val)
}

func (v *ArrayNode) Items() []JsonNode {
	if len(v.Objects) == 0 {
		return nil
	}
	var items []JsonNode
	for _, object := range v.Objects {
		node, err := wrap(object)
		if err != nil {
			continue
		}
		items = append(items, node)
	}
	return items
}

func asMapNode(value JsonNode) *MapNode {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case *MapNode:
		return value.(*MapNode)
	}
	return nil
}

func asArrayNode(value JsonNode) *ArrayNode {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case *ArrayNode:
		return value.(*ArrayNode)
	}
	return nil
}

func asValueNode(value JsonNode) *ValueNode {
	if value == nil {
		return nil
	}
	switch value.(type) {
	case *ValueNode:
		return value.(*ValueNode)
	}
	return nil
}

func (o *MapNode) Get(keys ...string) JsonNode {
	respMap := o.Object
	for i, key := range keys {
		value, exists := respMap[key]
		if !exists {
			return nil
		}
		if i == len(keys)-1 {
			node, err := wrap(value)
			if err != nil {
				return nil
			}
			return node
		}
		respMap = value.(map[string]interface{})
	}
	return nil
}

func ParseFileToMap(fp string) (*MapNode, error) {
	jsonNode, err := ParseFile(fp)
	if err != nil {
		return nil, err
	}
	switch jsonNode.(type) {
	case *MapNode:
		return jsonNode.(*MapNode), nil

	}
	return nil, errors.Errorf("The type is not a map %T", jsonNode)
}

func ParseFile(fp string) (JsonNode, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err //wrap
	}
	defer file.Close()
	return ParseReader(file)
}

func ParseReader(r io.Reader) (JsonNode, error) {
	jsonBytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err //wrap
	}
	return ParseBytes(jsonBytes)
}

func ParseBytes(b []byte) (JsonNode, error) {
	var in interface{}
	err := json.Unmarshal(b, &in)
	if err != nil {
		return nil, err //wrap
	}
	return wrap(in)
}

func wrap(in interface{}) (JsonNode, error) {
	if in == nil {
		return nil, errors.New("input is nil")
	}
	switch in.(type) {
	case []interface{}:
		return toArrayNode(in.([]interface{})), nil
	case map[string]interface{}:
		return toObjectNode(in.(map[string]interface{})), nil
	}
	return toValueNode(in), nil
}

func toValueNode(in interface{}) *ValueNode {
	return &ValueNode{Val: in}
}

func toArrayNode(items []interface{}) *ArrayNode {
	return &ArrayNode{
		Objects: items,
	}
}

func toObjectNode(in map[string]interface{}) *MapNode {
	return &MapNode{
		Object: in,
	}
}

func AsValue(n JsonNode) *ValueNode {
	return asValueNode(n)
}
func AsMap(n JsonNode) *MapNode {
	return asMapNode(n)
}
func AsArray(n JsonNode) *ArrayNode {
	return asArrayNode(n)
}
