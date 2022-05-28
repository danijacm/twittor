package main

import (
	"log"

	"github.com/danijacm/twittor/bd"
	"github.com/danijacm/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
