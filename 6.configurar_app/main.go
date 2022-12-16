package main

import (
	"configs/configuration"
	"configs/database"
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var id int
	ctx := context.Background()

	config, err := configuration.Load("config.yml")

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.CreateConnection(config)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRowContext(ctx, "SELECT id FROM books").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
	fmt.Println("Ok!")

}
