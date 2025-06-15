package main

import (
	"ip/server-search/app"
	"log"
	"os"
)

func main() {
	aplication := app.Generate()
	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
