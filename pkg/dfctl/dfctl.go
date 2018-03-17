package main

import (
	"fmt"
	"os"

	"github.com/yolo3301/dumb-flow/pkg/dfctl/cmd"
)

func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
