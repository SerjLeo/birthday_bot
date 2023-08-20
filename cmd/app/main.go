package main

import (
	"github.com/SerjLeo/birthday_bot/internal/app"
	"log"
)

func main() {
	err := app.InitApp()
	if err != nil {
		log.Fatal(err.Error())
	}
}
