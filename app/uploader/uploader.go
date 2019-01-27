package uploader

import (
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Uploader struct {
	Queue    []os.FileInfo
	Finished []string
	WorkFunc func(os.FileInfo)
	bar      *pb.ProgressBar
	mux      sync.Mutex
}

func (self *Uploader) Start() {
	self.Finished = []string{}
	var workers = 10
	wg.Add(workers)
	count := len(self.Queue)
	self.bar = pb.StartNew(count)
	for i := 1; i <= workers; i++ {
		go self.work()
	}
	wg.Wait()
	self.bar.FinishPrint("Finished uploading")
}

func (self *Uploader) work() {
	defer wg.Done()
	for {
		w, err := self.getWork()

		if err != nil {
			break
		}

		self.WorkFunc(w)
		self.bar.Increment()

	}
}

func (self *Uploader) getWork() (os.FileInfo, error) {
	self.mux.Lock()
	defer self.mux.Unlock()

	if len(self.Queue) == 0 {
		return nil, errors.New("Out of work")
	}

	x := self.Queue[0]

	if len(self.Queue) > 1 {
		self.Queue = self.Queue[1:]
	} else {
		self.Queue = []os.FileInfo{}
	}
	return x, nil
}

func GetContentType(path string) string {
	f, err := os.Open(path)
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return ""
	}
	f.Close()

	return http.DetectContentType(buffer)
}
func UploadToUrlWithRetries(url string, path string, retry int) {
	if url == "" {
		return
	}

	contentType := GetContentType(path)
	var b bytes.Buffer
	f, err := os.Open(path)

	if err != nil {
		fmt.Println(path + ": Could not open file")
		return
	}
	defer f.Close()

	var httpClient = &http.Client{
		Timeout: time.Second * 1000,
	}

	req, err := http.NewRequest("PUT", url, io.MultiReader(&b, f))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", contentType)
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if res.StatusCode != 200 {
		if retry == 0 {
			res.Write(os.Stderr)
			fmt.Println(err)
			return
		} else {
			UploadToUrlWithRetries(url, path, retry)
		}
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("bad status: %s", res.Status)
	}
	return
}
func UploadToUrl(url string, path string) {
	UploadToUrlWithRetries(url, path, 5)
}
