package models

import (
	"encoding/json"
	"log"
)

type (
	Annotation struct {
		Name `json:"name"`
		File string `json:"annotation_string"`
		Type string `json:"type"`
	}

	DatasetsAnnotation struct {
		Annotation `json:"annotation"`
	}

	DatasetsAnnotations struct {
		Data []DatasetsAnnotations
	}
)

func JsonToAnnotations(resp []byte) DatasetsAnnotations {
	var d DatasetsAnnotations

	err := json.Unmarshal(resp, &d)

	if err != nil {
		log.Fatal(err)
	}

	return d
}

// func (self DatasetsAnnotation) Save() {
// 	configJson, _ := json.Marshal(Store)
// 	err := ioutil.WriteFile(".labellab/files/", configJson, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	return
// }
