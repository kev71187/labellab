package main

import (
	"app/cloner"
	"app/config"
	"app/models"
	"app/puller"
	"app/views/prompts"
	"app/views/users"
	"errors"
	"fmt"
	"os"
	"strings"
)

var Env string

func getNames(inp string) ([]string, error) {
	resp := strings.Split(inp, "/")

	if len(resp) != 2 {
		return nil, errors.New("Format must be $ labellab username/dataset upload ./")
	}
	return resp, nil
}

func main() {
	// storage.Main()
	args := os.Args
	actionCount := len(args)

	if actionCount == 1 || args[0] == "help" {
		prompts.Help("base")
		return
	}
	arg1 := args[1]

	if arg1 == "login" {
		users.Login()
		return
	}

	if args[1] == "remote" {
		fmt.Println(config.BaseUrl)
		return
	}

	if arg1 == "auth" {
		users.Auth(args[2])
		return
	}
	empty := models.CurrentUser{}
	if models.Store.CurrentUser == empty {
		fmt.Println("You need to authenticate. Visit https://www.labellab.io/auth and paste the command shown into your command line below")
		return
	}
	names, err := getNames(args[1])

	if actionCount >= 1 && args[1] == "clone" {
		var clone_to string

		if actionCount >= 4 {
			clone_to = args[3]
		} else {
			clone_to = "./"
		}

		cloner.Clone(args[2], clone_to)
		return
	}

	if actionCount >= 1 && args[1] == "pull" {
		puller.Pull()
		return
	}

	if actionCount >= 3 && args[2] == "upload" {
		// dir, _ := filepath.Abs(filepath.Dir(args[0]))
		dir := args[3]
		if err != nil {
			fmt.Println(err)
			return
		}
		dataset := models.GetDatasetUsername(names[0], names[1])
		dataset.Upload(dir)
		// prompts.Help("dataset")
		return
	}

	if actionCount >= 2 {
		prompts.Help("dataset")
		return
	}

	// action := args[2]
	fmt.Println(args)
}
