package models

import (
	"app/uploader"
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
)

type (
	Dataset struct {
		Id              uint64 `json:"id"`
		CreatedAt       string `json:"created_at"`
		UserId          uint64 `json:"user_id"`
		Slug            string `json:"slug"`
		Username        string `json:"username"`
		AnnotationCount int    `json:"annotations_count"`
		DocumentCount   int    `json:"document_count"`
		Name            string `json:"name_slug"`
	}
	data struct {
		Data Dataset `json:"data"`
	}
	errorMessage struct {
		Message string `json:"message"`
	}
)

func (self Dataset) DatasetsDocumentsDeleted(start_date string, page int) DeletedDatasetsDocuments {
	url := "datasets_documents?basic=true&dataset_id=" + IdToString(self.Id) + "&limit=500&page=" + strconv.Itoa(page) + "&deleted=true&start_date=" + url.QueryEscape(start_date)
	resp := Request("GET", url, nil)
	r := JsonToDeletedDocuments(resp)
	return r
}

func (self Dataset) DatasetsDocuments(start uint64, page int) DatasetsDocuments {
	url := "datasets_documents?basic=true&dataset_id=" + IdToString(self.Id) + "&limit=500&page=" + strconv.Itoa(page) + "&start=" + IdToString(start)
	resp := Request("GET", url, nil)
	r := JsonToDocuments(resp)
	return r
}

func (self Dataset) GetDocumentAnnotations(start string, page int) DatasetsDocuments {
	url := "datasets_annotations?dataset_id=" + IdToString(self.Id) + "&limit=100&page=" + strconv.Itoa(page) + "&start=" + url.QueryEscape(start)
	resp := Request("GET", url, nil)
	r := JsonToDocuments(resp)
	return r
}

func (self Dataset) Upload(directory string) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		fmt.Println(err)
	}

	batch := CreateBatch(self.Id)

	up := uploader.Uploader{Queue: files, WorkFunc: func(file os.FileInfo) {
		UploadFile(directory, file, batch)
	}}

	up.Start()
}

func IdToString(id uint64) string {
	var r = strconv.FormatUint(id, 10)
	return r
}

func JsonToDataset(resp []byte) Dataset {
	var d data

	err := json.Unmarshal(resp, &d)
	if err != nil {
		log.Fatal(err)
	}

	return d.Data
}

func GetDatasetUsername(username string, name string) Dataset {
	resp := Request("GET", "datasets/"+username+"/"+name, nil)
	return JsonToDataset(resp)
}

func GetDataset(id uint64) Dataset {
	resp := Request("GET", "datasets/"+IdToString(id), nil)
	return JsonToDataset(resp)
}
