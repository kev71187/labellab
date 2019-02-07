package downloader

import (
	"app/config"
	"app/models"
	// "encoding/json"
	"errors"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type (
	Downloader struct {
		Queue    []models.DatasetsDocument
		Finished []string
		WorkFunc func(models.DatasetsDocument)
		Files    int
		bar      *pb.ProgressBar
		mux      sync.Mutex
		Config   config.LabelLabConfig
	}
)

func (self *Downloader) getWork() (models.DatasetsDocument, error) {
	self.mux.Lock()
	defer self.mux.Unlock()

	if len(self.Queue) == 0 {
		return models.DatasetsDocument{}, errors.New("Out of work")
	}
	x := self.Queue[0]
	if len(self.Queue) > 1 {
		self.Queue = self.Queue[1:]
	} else {
		self.Queue = []models.DatasetsDocument{}
	}

	return x, nil
}

func (self *Downloader) work() {
	defer wg.Done()
	for {
		if len(self.Queue) == 0 {
			break
		}
		w, err := self.getWork()

		if err != nil {
			break
		}
		self.WorkFunc(w)
		self.bar.Increment()

	}
}

func (self *Downloader) Start(start int) {
	self.Finished = []string{}
	self.bar = pb.StartNew(self.Files - start)
}

func (self *Downloader) Continue(queue []models.DatasetsDocument) {
	self.Queue = queue
	var workers = runtime.NumCPU() * 4
	wg.Add(workers)
	// count := len(self.Queue)
	for i := 1; i <= workers; i++ {
		go self.work()
	}
	wg.Wait()
}

func makeRange(min, max int) []string {
	a := make([]string, max-min+1)
	for i := range a {
		a[i] = string(min + i)
	}
	return a
}

func StringToId(id string) (uint64, error) {
	return strconv.ParseUint(id, 0, 64)
}

func getOffset(dir string) (int, uint64) {
	files, _ := ioutil.ReadDir(dir)
	var max = uint64(0)

	for _, file := range files {
		s := strings.Split(file.Name(), ".")
		if len(s) != 2 {
			continue
		}

		i, _ := StringToId(s[0])

		if i > max {
			max = i
		}
	}

	return len(files), max
}

func pullAnnotations(dataset models.Dataset, labelsDirectory string, conf config.LabelLabConfig) {
	var bar *pb.ProgressBar

	for i := 0; true; i++ {
		dd := dataset.GetDocumentAnnotations(conf.UpdatedAt, i)
		if i == 0 && dd.Pagination.Count > 0 {
			fmt.Println("Syncing annotations")
			bar = pb.StartNew(dd.Pagination.Count)
		}
		if len(dd.Data) == 0 {
			break
		}
		dd.SaveAnnotations(labelsDirectory)
		bar.Add(len(dd.Data))
	}
}

func deleteFiles(dataset models.Dataset, directory string, conf config.LabelLabConfig) {
	var bar *pb.ProgressBar

	for i := 0; true; i++ {
		dd := dataset.DatasetsDocumentsDeleted(conf.UpdatedAt, i)
		if i == 0 && dd.Pagination.DocumentCount > 0 {
			fmt.Println("Deleting files:")
			bar = pb.StartNew(dd.Pagination.DocumentCount)
		}
		if len(dd.Data) == 0 {
			break
		}
		for j := range dd.Data {
			var doc = dd.Data[j]
			doc.Destroy(directory)
			bar.Add(1)
		}
		// dd.SaveAnnotations(labelsDirectory)
	}
}

func Download(dataset models.Dataset, directory string, conf config.LabelLabConfig) int {
	filesDirectory := directory + "/files"
	labelsDirectory := directory + "/labels"
	timer := time.Now()

	deleteFiles(dataset, directory, conf)

	down := Downloader{Config: conf, Files: dataset.DocumentCount, WorkFunc: func(dataset_document models.DatasetsDocument) {
		models.DownloadFile(filesDirectory, dataset_document)
	}}

	var totalStart, start = getOffset(filesDirectory)
	var totalToDownload = down.Files - totalStart
	if totalToDownload > 0 {
		fmt.Println("Downloading:", conf.Dataset, "(", dataset.DocumentCount, "Total files ) (", totalToDownload, " Remaining to download )")
		down.Start(totalStart)

		for i := 0; true; i++ {
			dd := dataset.DatasetsDocuments(start, i)

			if len(dd.Data) == 0 {
				break
			}
			down.Continue(dd.Data)
		}

		down.bar.FinishPrint("Finished downloading files")
	}

	pullAnnotations(dataset, labelsDirectory, conf)
	conf.Touch(timer, directory)
	fmt.Println("Pull complete")
	return 1
}
