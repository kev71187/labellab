package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"
)

type (
	LabelLabConfig struct {
		Version     int    `json:"version"`
		Files       int    `json:"files"`
		Annotations int    `json:"annotations"`
		Dataset     string `json:"dataset"`
		UpdatedAt   string `json:"updated_at"`
	}
)

func (self LabelLabConfig) Touch(at time.Time, directory string) {
	self.UpdatedAt = at.Format(time.RFC850)
	self.Save(directory)
}

func (self LabelLabConfig) Save(dir string) {
	confJson, _ := json.Marshal(self)
	err := ioutil.WriteFile(dir+"/dataset.json", jsonPrettyPrint(confJson), 0644)
	if err != nil {
		panic(err)
	}
}

func jsonPrettyPrint(in []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, in, "", "  ")
	if err != nil {
		return in
	}
	return out.Bytes()
}
