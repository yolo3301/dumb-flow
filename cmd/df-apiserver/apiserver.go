package main

import (
	"log"
	"os"

	"github.com/yolo3301/dumb-flow/cmd/df-apiserver/app"
)

func main() {
	table, err := app.NewDefaultTableDAO()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer table.Close()

	queue := app.DefaultQueueDAO{}

	server := app.NewDumbflowServer(table, queue)
	server.Run()
}
