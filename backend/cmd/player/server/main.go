package main

import (
	"github.com/sky0621/my-test-project/backend/player/setup"
	"log"
)

func main() {
	if err := setup.NewApp().Run(); err != nil {
		log.Fatal(err)
	}
}
