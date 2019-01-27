package puller

import (
	"app/downloader"
	"app/loader"
)

func Pull() error {
	conf, err := loader.LoadLabelLabConfig()

	if err != nil {
		return err
	}

	dataset := loader.LoadDataset(conf)
	downloader.Download(dataset, ".", conf)
	return nil
}
