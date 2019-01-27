package cloner

import (
	"app/downloader"
	"app/loader"
	"app/models"
	"fmt"
	"os"
	"strings"
)

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func Clone(dataset_name string, directory string) int {
	s := strings.Split(dataset_name, "/")
	suf := strings.HasSuffix(directory, "/")
	if !suf {
		directory = directory + "/"
	}

	var rootDir = directory + s[1]
	dataset := models.GetDatasetUsername(s[0], s[1])
	if Exists(rootDir) {
		fmt.Println("Cannot clone to this location. A folder with the name \"" + dataset.Name + "\" already exists in this location. \n  If this is a labellab dataset you can change into that directory and run: \n    $ labellab pull")
		return 1
	}
	fmt.Println("Cloning " + dataset.Name)
	conf := loader.NewLLConfig(dataset)
	var filesDir = rootDir + "/files"
	os.Mkdir(rootDir, os.FileMode(0744))
	os.Mkdir(rootDir+"/labels", os.FileMode(0744))
	os.Mkdir(filesDir, os.FileMode(0744))

	conf.Save(rootDir)
	fmt.Println(conf)

	downloader.Download(dataset, rootDir, conf)
	return 1
}
