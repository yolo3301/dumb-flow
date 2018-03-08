package main

import "github.com/yolo3301/dumb-flow/cmd/df-apiserver/app"

func main() {
	table := app.NewDefaultTableDAO()
	queue := app.DefaultQueueDAO{}

	server := app.NewDumbflowServer(table, queue)
	server.Run()
}
