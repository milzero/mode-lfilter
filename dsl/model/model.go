package model

import (
	"fmt"
	"sync"
)

type Meta interface {
	ParseFrom(map[string]interface{})
}

type MetaFromJson struct {
	mtx  sync.Mutex
	meta map[string]interface{}
}

func (m *MetaFromJson) ParseFrom(inpiut map[string]interface{}) {
	m.renew()
	var value interface{}
	meta := map[string]interface{}{}
	for k, v := range inpiut {
		if _, ok := m.meta[k]; ok {
			switch val := i.(type) {
			case int:
				fmt.Printf("Twice %v is %v\n", v, val*2)
			case string:
				fmt.Printf("%q is %v bytes long\n", v, len(val))
			default:
				fmt.Printf("I don't know about type %T!\n", val)
			}
		}
	}
}

func (m *MetaFromJson) ParseFromRecursion(inpiut map[string]interface{}) {

}

func (m MetaFromJson) renew() {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.meta = nil
	m.meta = map[string]interface{}{}
}
