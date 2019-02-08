package models

import (
	"app/uploader"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	Document struct {
		Id        uint64 `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		UserId    uint64 `json:"user_id"`
		File      File   `json:"file"`
	}

	File struct {
		Url string `json:"url"`
	}

	Name struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	// Annotation struct {
	// 	Name `json:"name"`
	// 	File string `json:"annotation_string"`
	// 	Type string `json:"type"`
	// }

	// DatasetsAnnotation struct {
	// 	Annotation `json:"annotation"`
	// }

	DeletedDatasetsDocuments struct {
		Data       []DatasetsDocument
		Pagination pagination `json:"pagination"`
	}

	DatasetsDocument struct {
		DatasetsAnnotations []DatasetsAnnotation `json:"datasets_annotations"`
		Document            Document             `json:"document"`
		Id                  int                  `json:"id"`
		Annotations         string               `json:"annotations_string"`
		DeletedAt           string               `json:"deleted_at"`
	}

	DatasetsDocuments struct {
		Data       []DatasetsDocument
		Pagination pagination `json:"pagination"`
	}

	pagination struct {
		Count         int    `json:"count"`
		DocumentCount int    `json:"document_count"`
		Annotations   uint64 `json:"annotations"`
		PerPage       int    `json:"per_page"`
	}

	FileMetadata struct {
		Name        string `json:"file_name"`
		DatasetId   uint64 `json:"dataset_id"`
		BatchId     uint64 `json:"batch_id"`
		FileMd5     string `json:"content_md5"`
		ContentType string `json:"content_type"`
	}

	UploadFileUrl struct {
		Id        uint64 `json:"id"`
		UploadUrl string `json:"upload_url"`
	}

	uploadData struct {
		Data UploadFileUrl `json:"data"`
	}

	fileData struct {
		Data FileMetadata `json:"data"`
	}
)

func GetMD5Hash(bytes []byte) string {
	hasher := md5.New()
	hasher.Write(bytes)
	return hex.EncodeToString(hasher.Sum(nil))
}

func FileTypeMatch(name string) bool {
	s := strings.Split(name, ".")
	ext := s[len(s)-1]
	matched, _ := regexp.MatchString("jpg|jpeg|png|gif|csv|txt|json", strings.ToLower(ext))
	return matched
}

func JsonToDeletedDocuments(resp []byte) DeletedDatasetsDocuments {
	var d DeletedDatasetsDocuments

	err := json.Unmarshal(resp, &d)

	if err != nil {
		log.Fatal(err)
	}

	return d
}

func JsonToDocuments(resp []byte) DatasetsDocuments {
	var d DatasetsDocuments

	err := json.Unmarshal(resp, &d)

	if err != nil {
		log.Fatal(err)
	}

	return d
}

func (self DatasetsDocument) FileName() string {
	return strconv.Itoa(self.Id)
}
func (self DatasetsDocument) AnnotationsFileName() string {
	return strconv.Itoa(self.Id) + ".json"
}
func (self DatasetsDocument) DocumentFileName() string {
	s := strings.Split(self.Document.Name, ".")
	return self.FileName() + "." + s[1]
}

func (self DatasetsDocument) Destroy(directory string) {
	filePath := directory + "/files" + "/" + self.DocumentFileName()
	labelPath := directory + "/labels" + "/" + self.AnnotationsFileName()
	// labelsDirectory := directory + "/labels"
	if Exists(filePath) {
		os.Remove(filePath)
	}
	if Exists(labelPath) {
		os.Remove(labelPath)
	}
}
func (self DatasetsDocuments) SaveAnnotations(directory string) {
	for i := 0; i < len(self.Data); i++ {
		dd := self.Data[i]
		dd.SaveAnnotations(directory)
	}
}

func (self DatasetsDocument) SaveAnnotations(directory string) {
	var fileName = directory + "/" + self.AnnotationsFileName()
	err := ioutil.WriteFile(fileName, []byte(self.Annotations), 0644)

	if err != nil {
		panic(err)
	}
}

func saveUrlToFile(name string, url string) error {
	out, err := os.Create(name)
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func DownloadFile(directory string, dd DatasetsDocument) error {

	var fileName = directory + "/" + dd.DocumentFileName()

	if !Exists(fileName) {
		err := saveUrlToFile(fileName, dd.Document.File.Url)
		if err != nil {
			return err
		}
	}

	return nil
	// fmt.Println("Hi")
}

func UploadFile(directory string, file os.FileInfo, batch Batch) {
	name := file.Name()
	path := directory + "/" + name
	if FileTypeMatch(name) {
		f, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Println(err)
		}

		hash := GetMD5Hash(f)
		contentType := uploader.GetContentType(path)

		reqBody := fileData{
			Data: FileMetadata{Name: file.Name(),
				FileMd5:     hash,
				DatasetId:   batch.DatasetId,
				BatchId:     batch.Id,
				ContentType: contentType}}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(reqBody)
		resp := RequestWithRetries("POST", "documents", b, 5)
		var u uploadData
		err = json.Unmarshal(resp, &u)
		url := u.Data.UploadUrl
		if url == "" {
			return
		}
		uploader.UploadToUrl(url, path)
		RequestWithRetries("PUT", "documents/"+strconv.FormatUint(u.Data.Id, 10)+"/complete", nil, 5)
	}
}
