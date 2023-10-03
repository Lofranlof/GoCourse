//go:build !change

package main

import (
	"flag"

	"gitlab.com/manytask/itmo-go/public/coverme/app"
	"gitlab.com/manytask/itmo-go/public/coverme/models"
)

func main() {
	port := flag.Int("port", 8080, "port to listen")
	flag.Parse()

	db := models.NewInMemoryStorage()
	app.New(db).Start(*port)
}
