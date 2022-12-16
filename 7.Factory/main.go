package main

import (
	"factory/configuration"
	"factory/repository"
	"fmt"
	"log"
)

// Factory Pattern

func main() {
	config := &configuration.Configuration{
		Engine: "mysql",
		Host:   "localhost",
		User:   "golang",
		DBName: "books",
	}

	repo, err := repository.New(config)
	if err != nil {
		log.Fatal(err)
	}

	err = repo.Save("")
	if err != nil {
		log.Fatal(err)
	}

	data := repo.Find(1)
	fmt.Println(data)

}
