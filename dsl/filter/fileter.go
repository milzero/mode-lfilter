package filter

import (
	"fmt"
	"reflect"

	"github.com/milzero/mode-lfilter/dsl/model"
	"go.uber.org/zap"
)

type Type int

func Init() {
	InitLog()
}

func InitLog() {
	zap.L().Core().Enabled(zap.DebugLevel)
}

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

	logger *zap.Logger

	priorityFilter struct {
		priorityIndex []int
		filters       map[int]*Filter
	}
}

func (model *Model) createPriorityFilter() error {

	if model.logger == nil {
		model.logger = zap.L()
	}

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

func (model *Model) process(meta model.Meta) (map[string]interface{}, error) {

	metaDict := meta.GetMeta()
	profile := map[string]interface{}{}
	c := newCompare()

	for _, idx := range model.priorityFilter.priorityIndex {
		filter, ok := model.priorityFilter.filters[idx]
		if !ok {
			model.logger.Warn("priority filter not found", zap.Any("idx", idx))
			continue
		}

		methodName := filter.Method
		c.isFunc(methodName)
		m, ok := metaDict[filter.Type]
		if !ok {
			continue
		}

		for _, item := range filter.Items {
			in := make([]reflect.Value, 2)
			in[0] = reflect.ValueOf(m)
			in[1] = reflect.ValueOf(item.Key)
			out := reflect.ValueOf(c).MethodByName(methodName).Call(in)
			if len(out) > 0 {
				ret := out[0].Interface()
				b, ok := ret.(bool)
				if ok {
					if b {
						profile[filter.Type] = item.Value
					}
				}
			}

		}
	}
	return profile, nil
}
