package models

import (
	"encoding/json"
	"log"
)

type (
	dataBatch struct {
		Data Batch `json:"data"`
	}
	Batch struct {
		Id        uint64 `json:"id"`
		DatasetId uint64 `json:"dataset_id"`
	}
)

func JsonToBatch(resp []byte) Batch {
	var d dataBatch

	err := json.Unmarshal(resp, &d)
	if err != nil {
		log.Fatal(err)
	}

	return d.Data
}

func CreateBatch(id uint64) Batch {
	resp := Request("POST", "batches?dataset_id="+IdToString(id), nil)
	var d dataBatch
	err := json.Unmarshal(resp, &d)
	if err != nil {
		log.Fatal(err)
	}
	return d.Data
}
