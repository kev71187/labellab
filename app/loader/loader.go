package loader

import (
	. "app/config"
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func LoadDataset(conf LabelLabConfig) models.Dataset {
	s := strings.Split(conf.Dataset, "/")
	return models.GetDatasetUsername(s[0], s[1])
}
func LoadLabelLabConfig() (LabelLabConfig, error) {
	var fileName = "/dataset.json"
	pwd, _ := os.Getwd()
	var path = pwd + fileName

	if !Exists(path) {
		var e = "No dataset.json located in this directory"
		log.Fatal(e)
		return LabelLabConfig{}, errors.New(e)
	}

	txt, _ := ioutil.ReadFile(path)
	var c LabelLabConfig

	err := json.Unmarshal(txt, &c)
	if err != nil {
		log.Fatal(err)
		return LabelLabConfig{}, err
	}

	return c, nil

}

func NewLLConfig(dataset models.Dataset) LabelLabConfig {
	return LabelLabConfig{
		Version:     1,
		Dataset:     trimFirstRune(dataset.Slug),
		Files:       dataset.DocumentCount,
		Annotations: dataset.AnnotationCount,
	}
}
