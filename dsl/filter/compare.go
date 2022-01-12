package filter

import (
	"log"
	"reflect"
	"regexp"
	"strings"
)

type compare struct {
	methods map[string]struct{}
}

func newCompare() *compare {
	return &compare{}
}

func (c *compare) Match(k1, k2 interface{}) bool {
	switch k1.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string:
		if reflect.TypeOf(k1) == reflect.TypeOf(k2) {
			return k1 == k2
		}
	default:
		log.Println("default")
	}
	return false
}

func (c *compare) Contain(k1, k2 string) bool {
	return strings.Contains(k1, k2)
}

func (c *compare) Regex(regex, k2 string) bool {
	match, err := regexp.MatchString(regex, k2)
	if err != nil {
		return false
	}
	return match
}

func (c *compare) Great(k1, k2 interface{}) bool {
	switch k1.(type) {
	case int:
		kv, ok := k2.(int)
		if !ok {
			return false
		}
		return k1.(int) > kv
	case float64:
		kv, ok := k2.(float64)
		if !ok {
			return false
		}
		return k1.(float64) > kv
	}
	return false
}

func (c *compare) Equal(k1, k2 interface{}) bool {
	return k1 == k2
}
func (c *compare) Less(k1, k2 interface{}) bool {
	switch k1.(type) {
	case int:
		kv, ok := k2.(int)
		if !ok {
			return false
		}
		return k1.(int) < kv
	case float64:
		kv, ok := k2.(float64)
		if !ok {
			return false
		}
		return k1.(float64) < kv
	}
	return false
}

func (c *compare) isFunc(method string) bool {

	if c.methods == nil {
		c.methods = map[string]struct{}{}
		value := reflect.ValueOf(c)
		num := value.NumMethod()
		for i := 0; i < num; i++ {
			methodName := value.Method(i).String()
			c.methods[methodName] = struct{}{}
		}
	}

	_, ok := c.methods[method]
	return ok
}
