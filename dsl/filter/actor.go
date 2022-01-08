package filter

import (
	"fmt"
	"github.com/milzero/mode-lfilter/dsl/model"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"sync"
)

type Actor struct {
	modelContainer *ModelContainer
	isInit         bool
	mtx            sync.Mutex
	logger         *zap.Logger
}

func NewActuator() *Actor {
	return &Actor{
		modelContainer: &ModelContainer{},
		isInit:         false,
		mtx:            sync.Mutex{},
		logger:         zap.L(),
	}
}

func (a *Actor) Precess(meta model.Meta) error {
	a.modelContainer.FilterModel.process(meta)
	return nil
}

func (a *Actor) Init(bytes []byte, structType Type) error {
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

func (a *Actor) parseYaml(bytes []byte) error {
	modelContainer := &ModelContainer{}
	err := yaml.Unmarshal(bytes, modelContainer)
	if err != nil {
		return err
	}

	a.modelContainer = modelContainer
	return nil
}

func (a *Actor) check(container *ModelContainer) error {
	if container == nil {
		return fmt.Errorf("input is nil")
	}

	if container.FilterModel == nil {
		a.logger.Error("input filterModel is nil")
		return fmt.Errorf("input filterModel is nil")
	}

	if container.FilterModel.Filters == nil {
		a.logger.Error("input filters is nil")
		return fmt.Errorf("input filters is nil")
	}

	filterLen := len(container.FilterModel.Filters)
	if filterLen < 1 {
		a.logger.Error("model container lenght is 0")
		return fmt.Errorf("modelContainer  lenght is 0")
	}

	filters := make(map[int]*Filter)
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

func (a *Actor) UpdateFilter(bytes []byte) error {

	a.logger.Debug("Actor filter updated")
	err := a.parseYaml(bytes)
	if err != nil {
		a.logger.Error("parse incoming updated error", zap.String("error", err.Error()))
		return err
	}

	a.logger.Debug("Actor filter updated ")
	return nil
}
