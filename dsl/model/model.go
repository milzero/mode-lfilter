package model

import (
	"sync"
)

type Meta interface {
	ParseFrom(map[string]interface{})
}

type MetaFromJson struct {
	mtx  sync.Mutex
	meta map[string]interface{}
}

func (m *MetaFromJson) ParseFrom(dict map[string]interface{}) {
	m.renew()
	meta := map[string]interface{}{}
	for k, v := range dict {
		if _, ok := m.meta[k]; !ok {
			meta[k] = v
		}
	}
}

func (m MetaFromJson) renew() {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.meta = nil
	m.meta = map[string]interface{}{}
}
