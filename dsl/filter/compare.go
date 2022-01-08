package filter

import (
	"regexp"
	"strings"
)

type compare struct {
}

func newCompare() *compare {
	return &compare{}
}

func (c *compare) match(k1, k2 string) bool {
	//string k1, k2 are equal
	return k1 == k2
}

func (c *compare) contain(k1, k2 string) bool {
	//string k1 contains k2
	return strings.Contains(k1, k2)
}

func (c *compare) regex(regex, k2 string) bool {
	match, err := regexp.MatchString(`^Golang`, k2)
	if err != nil {
		return false
	}
	return match
}

func (c *compare) great(k1, k2 interface{}) bool {
	//TODO 大于
	return false
}

func (c *compare) equal(k1, k2 interface{}) bool {
	//TODO 等于
	return false
}
func (c *compare) less(k1, k2 interface{}) bool {
	//TODO 小于
	return false
}

var c compare
