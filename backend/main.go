package main

import (
	"backend/loaders/fiber"
	"backend/loaders/fs"
)

func main() {
	fs.Init()
	fiber.Init()
}
