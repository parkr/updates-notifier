package main

import (
	"log"
	"os"

	updatesnotifier "github.com/parkr/updates-notifier"
	"github.com/parkr/updates-notifier/docker"
)

func must(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %+v", message, err)
	}
}

func main() {
	config := updatesnotifier.Config{
		DockerRepositories: []docker.Repository{
			{
				Name: "parkr/radar",
				Tag:  "42b7a780170921975730b1c12c4b865f0de45303",
			},
		},
	}
	errs := updatesnotifier.Run(config)
	if len(errs) == 0 {
		log.Println("No updates!")
		return
	}
	for _, err := range errs {
		log.Println(err)
	}
	os.Exit(1)
}
