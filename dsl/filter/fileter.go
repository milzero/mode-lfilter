package filter

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"sync"
)

type Type int

const (
	YAML = iota
	JSON = iota + 1
	XML  = iota + 2
)

type ModelContainer struct {
	FilterModel Model `yaml:"filterModel" json:"filterModel" xml:"filterModel"`
}
type Items struct {
	Key   string `yaml:"key" json:"key,omitempty" xml:"key"`
	Value string `yaml:"value" json:"value,omitempty" xml:"value"`
}
type Filters struct {
	Desc     string  `yaml:"desc" json:"desc" json:"desc,omitempty" xml:"desc"`
	Function string  `yaml:"function" json:"function" json:"function,omitempty" xml:"function"`
	Type     string  `yaml:"type" json:"type" json:"type,omitempty" xml:"type"`
	Method   string  `yaml:"method" json:"method,omitempty" xml:"method"`
	Items    []Items `yaml:"items" json:"items,omitempty" xml:"items"`
}
type Model struct {
	Namespace string    `yaml:"namespace" json:"namespace,omitempty" xml:"namespace"`
	Desc      string    `yaml:"desc" json:"desc,omitempty" xml:"desc"`
	Filters   []Filters `yaml:"filtes" json:"filters,omitempty" xml:"filters"`
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
