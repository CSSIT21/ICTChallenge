package main

import (
	"backend/loaders/fiber"
	"backend/loaders/hub"
)

func main() {
	hub.Init()
	fiber.Init()
}
