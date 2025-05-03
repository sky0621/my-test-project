package main

import (
	"github.com/sky0621/my-test-project/app"
	"log"
)

func main() {
	if err := app.New().Run(); err != nil {
		log.Fatal(err)
	}
}
