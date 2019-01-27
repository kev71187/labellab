package fixtures

import (
	"fmt"
	"io/ioutil"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DefaultDownloadConfig() string {
	dat, err := ioutil.ReadFile("/go/src/app/fixtures/data/downloader/default_labellab_config.json")
	fmt.Println(time.Now().Format(time.RFC850))

	check(err)
	return string(dat)
}
