//go:build !change

package main

import (
	"flag"

	"gitlab.com/manytask/itmo-go/private/coverme/app"
	"gitlab.com/manytask/itmo-go/private/coverme/models"
)

func main() {
	port := flag.Int("port", 8080, "port to listen")
	flag.Parse()

	db := models.NewInMemoryStorage()
	app.New(db).Start(*port)
}
