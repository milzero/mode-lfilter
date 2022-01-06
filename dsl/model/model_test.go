package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func TestMetaFromJson_ParseFrom(t *testing.T) {
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

	var mm MetaFromJson
	mm.ParseFrom(m)
	mm.ParseFrom(m)
}
