package cmd

import (
	"errors"
	"os"

	"faceit_users/cmd/consumer"
	"faceit_users/cmd/producer"
	"faceit_users/cmd/usersapi"
)

func Execute() error {
	apps := initializeApplications()

	currentApp := os.Getenv("APP")

	exsitApp := apps[currentApp]

	if exsitApp == nil {
		return errors.New("Application not found")
	}

	return exsitApp()
}

func initializeApplications() map[string]func() error {
	dict := make(map[string]func() error)

	dict["consumer"] = func() error {
		return consumer.Execute(os.Args)
	}
	dict["userapi"] = func() error {
		return usersapi.Execute(os.Args)
	}
	dict["producer"] = func() error {
		return producer.Execute(os.Args)
	}
	return dict

}
