package prompts

import (
	"fmt"
)

func Logo() string {
	return `
Label Lab Io - Help
`
}

func Help(instructions string) {
	var options = ""
	if instructions == "dataset" {
		options = Logo() + `
  Dataset Options:
      pull                      - downloads/syncs the files and labels stored on label lab to your device. Must be in the root directory of your dataset
                                    $ labellab pull

      upload                    - file or directory to upload
                                    $ labellab example-user/faces upload { directory_to_upload }
`
	} else if instructions == "base" {
		options = Logo() + `
  Base Options:
      auth                      - login to your open images account
                                    $ labellab auth { auth_hash_found_at } https://www.labellab.io/auth

      clone                      - clone your dataset to your device
                                    $ labellab clone { username/dataset }

      pull                      - downloads/syncs the files and labels stored on label lab to your device. Must be in the root directory of your dataset
                                    $ labellab pull

      upload                    - file or directory to upload
                                    $ labellab example-user/faces upload { directory_to_upload }
`
	}
	fmt.Println(options)
}
