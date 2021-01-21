package main

import (
	"log"

	"github.com/Oscar02-94/twittor/bd"
	"github.com/Oscar02-94/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion ala bd")
		return
	}
	handlers.Manejadores()
}
