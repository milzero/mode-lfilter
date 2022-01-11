package filter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
	"testing"

	"github.com/milzero/mode-lfilter/dsl/model"

	"gopkg.in/yaml.v2"
)

func TestName(t *testing.T) {
	conf := &ModelContainer{}
	yamlFile, err := ioutil.ReadFile("filters.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
}

func TestActuator_Init(t *testing.T) {
	actuator := Actor{
		modelContainer: &ModelContainer{},
		isInit:         false,
		mtx:            sync.Mutex{},
	}

	yamlBytes, err := ioutil.ReadFile("filters.yaml")
	if err != nil {
		log.Printf("read yaml faile: %s", err)
	}

	actuator.Init(yamlBytes, YAML)
}

func TestActuator_Init1(t *testing.T) {
	actuator := Actor{
		modelContainer: &ModelContainer{},
		isInit:         false,
		mtx:            sync.Mutex{},
	}

	yamlBytes, err := ioutil.ReadFile("filters.yaml")
	if err != nil {
		log.Printf("read yaml faile: %s", err)
	}

	actuator.Init(yamlBytes, YAML)
}

func TestActuator_Precess(t *testing.T) {
	bytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Panic("read file failed")
	}

	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		log.Panic("unmarshal failed", err)
	}

	log.Printf("unmarshal %+v", m)

	var mm model.Meta = model.NewMetaFromJson()
	mm.ParseFrom(m)

	a := NewActuator()

	bytes, err = ioutil.ReadFile("filters.yaml")
	if err != nil {
		log.Panic("read file failed")
	}
	a.Init(bytes, YAML)
	metas, er := a.Process(mm)
	if err != nil {
		log.Panic(er)
	}

	bs, err := json.Marshal(metas)
	if err != nil {
		log.Panic(er)
	}

	log.Println(string(bs[:]))
}
