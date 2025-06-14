package main

import (
	"fmt"
	"ip/server-search/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start point")

	aplication := app.Generate()
	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
