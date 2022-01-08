package filter

import (
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
	//string k1, k2 are equal
	return k1 == k2
}

func (c *compare) Contain(k1, k2 string) bool {
	//string k1 contains k2
	return strings.Contains(k1, k2)
}

func (c *compare) Regex(regex, k2 string) bool {
	match, err := regexp.MatchString(`^Golang`, k2)
	if err != nil {
		return false
	}
	return match
}

func (c *compare) Great(k1, k2 interface{}) bool {
	//TODO 大于
	return false
}

func (c *compare) Equal(k1, k2 interface{}) bool {
	//TODO 等于
	return false
}
func (c *compare) Less(k1, k2 interface{}) bool {
	//TODO 小于
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
