package model

import (
	"sync"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type Meta interface {
	ParseFrom(map[string]interface{})
	GetMeta() map[string]interface{}
}

type MetaFromJson struct {
	mtx    sync.Mutex
	meta   map[string]interface{}
	logger *zap.Logger
}

func NewMetaFromJson() *MetaFromJson {
	return &MetaFromJson{
		mtx:    sync.Mutex{},
		meta:   map[string]interface{}{},
		logger: zap.L(),
	}
}

func (m *MetaFromJson) ParseFrom(dict map[string]interface{}) {
	m.logger.Debug("starting to parse dick", zap.Any("dict", dict))
	m.renew()
	meta := map[string]interface{}{}
	for k, v := range dict {
		if _, ok := m.meta[k]; !ok {
			meta[k] = v
		}
	}
	m.meta = meta
	m.logger.Debug("finish to parse dick")
}

func (m MetaFromJson) renew() {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.meta = nil
	m.meta = map[string]interface{}{}
	m.logger.Debug("renew")
}

func (m MetaFromJson) GetMeta() map[string]interface{} {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	meta := map[string]interface{}{}
	copier.Copy(&meta, &m.meta)

	return meta
}
