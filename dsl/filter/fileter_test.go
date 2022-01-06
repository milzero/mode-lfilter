package filter

import (
	"io/ioutil"
	"log"
	"sync"
	"testing"

	yaml "gopkg.in/yaml.v2"
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
	actuator := Actuator{
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
