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
	FilterModel *Model `yaml:"filterModel" json:"filterModel" xml:"filterModel"`
}
type Item struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
type Filters struct {
	Desc     string  `yaml:"desc"`
	Function string  `yaml:"function"`
	Type     string  `yaml:"type"`
	Method   string  `yaml:"method"`
	Priority int     `yaml:"priority"`
	Items    []*Item `yaml:"items"`
}
type Model struct {
	Version   string     `yaml:"version"`
	Namespace string     `yaml:"namespace"`
	Desc      string     `yaml:"desc"`
	Filters   []*Filters `yaml:"filters"`

	priorityFilter struct {
		priorityIndex []int
		filters       map[int]*Filters
	}
}

func (model *Model) createPriorityFilter() error {
	if model.Filters == nil {
		return fmt.Errorf("Model.Filters is nil")
	}

	for _, filter := range model.Filters {
		if filter == nil {
			continue
		}
		idx := filter.Priority
		model.priorityFilter.priorityIndex = append(model.priorityFilter.priorityIndex, idx)
		model.priorityFilter.filters[idx] = filter
	}
	return nil
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

	err = a.modelContainer.FilterModel.createPriorityFilter()
	if err != nil {
		return err
	}

	return nil
}

func (a *Actuator) parseYaml(bytes []byte) error {
	modelContainer := &ModelContainer{}
	err := yaml.Unmarshal(bytes, modelContainer)
	if err != nil {
		return err
	}
	return nil
}

func (a *Actuator) check(container *ModelContainer) error {
	if container == nil {
		return fmt.Errorf("input is nil")
	}

	if container.FilterModel == nil {
		return fmt.Errorf("input filterModel is nil")
	}

	if container.FilterModel.Filters == nil {
		return fmt.Errorf("input filters is nil")
	}

	filterLen := len(container.FilterModel.Filters)
	if filterLen < 1 {
		return fmt.Errorf("modelContainer  lenght is 0")
	}

	filters := make(map[int]*Filters)
	var priority []int
	for _, filter := range container.FilterModel.Filters {
		if filter == nil {
			continue
		}
		filters[filter.Priority] = filter
		priority = append(priority, filter.Priority)
	}

	if len(filters) != filterLen {
		fmt.Errorf("duplicate filter priority")
	}

	err := a.modelContainer.FilterModel.createPriorityFilter()
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
