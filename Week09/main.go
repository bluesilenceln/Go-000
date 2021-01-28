package main

import (
	"Go-000/Week09/echo"
	"Go-000/Week09/pkg"
	"log"
)

func main() {
	s := echo.NewServer(":8080")

	app := pkg.New()
	app.Append(s)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		log.Printf("start failed: %v\n", err)
	}
}
