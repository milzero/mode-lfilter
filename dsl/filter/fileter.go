package filter

import (
	"fmt"
	"github.com/milzero/mode-lfilter/dsl/model"
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

const (
	Match   string = "match"
	Contain string = "contain"
	Regex   string = "regex"
)

type Filter struct {
	Desc     string  `yaml:"desc"`
	Function string  `yaml:"function"`
	Type     string  `yaml:"type"`
	Method   string  `yaml:"method"`
	Priority int     `yaml:"priority"`
	Items    []*Item `yaml:"items"`
}

type Model struct {
	Version   string    `yaml:"version"`
	Namespace string    `yaml:"namespace"`
	Desc      string    `yaml:"desc"`
	Filters   []*Filter `yaml:"filters"`

	priorityFilter struct {
		priorityIndex []int
		filters       map[int]*Filter
	}
}

func (model *Model) createPriorityFilter() error {
	if model.Filters == nil {
		return fmt.Errorf("Model.Filters is nil")
	}

	if model.priorityFilter.filters == nil {
		model.priorityFilter.filters = map[int]*Filter{}
		model.priorityFilter.priorityIndex = []int{}
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

func (model *Model) process(meta model.Meta) error {

	metaDict := meta.GetMeta()
	for i, _ := range model.priorityFilter.priorityIndex {
		filter, ok := model.priorityFilter.filters[i]
		if !ok {
			continue
		}

		m, ok := metaDict[filter.Type]
		if ok {
			fmt.Printf("%+v", m)
		}
	}
	return nil
}
