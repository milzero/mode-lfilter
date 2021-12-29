package filter

import (
	"fmt"
	"sync"

	"gopkg.in/yaml.v2"
)

type Type int

const (
	YAML = iota
	JSON = iota + 1
	XML  = iota + 2
)

type ModelContainer struct {
	FilterModel FilterModel `yaml:"filterModel" json:"filterModel" xml:"filterModel"`
}
type Items struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
type Filters struct {
	Desc     string  `yaml:"desc"`
	Function string  `yaml:"function"`
	Type     string  `yaml:"type"`
	Method   string  `yaml:"method"`
	Priority int     `yaml:"priority"`
	Items    []Items `yaml:"items"`
}
type FilterModel struct {
	Version   string    `yaml:"version"`
	Namespace string    `yaml:"namespace"`
	Desc      string    `yaml:"desc"`
	Filters   []Filters `yaml:"filters"`
}

type Actuator struct {
	modelContainer *ModelContainer
	isInit         bool
	mtx            sync.Mutex
}

func (a *Actuator) Init(bytes []byte, structType Type) error {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	a.isInit = false
	defer func() {
		a.isInit = true
	}()

	if a.isInit {
		return fmt.Errorf("filter have init")
	}

	var err error
	switch structType {
	case YAML:
		err = a.parseYaml(bytes)
	default:
		err = fmt.Errorf("unkown type")
	}

	if err != nil {
		return err
	}

	return nil
}

func (a *Actuator) parseYaml(bytes []byte) error {
	a.modelContainer = &ModelContainer{}
	err := yaml.Unmarshal(bytes, a.modelContainer)
	if err != nil {
		return err
	}
	return nil
}

func (a *Actuator) UpdateFilter(bytes []byte) error {
	err := a.parseYaml(bytes)
	if err != nil {
		return err
	}
	return nil
}
