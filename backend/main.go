package main

import (
	"backend/loaders/fiber"
	"backend/loaders/fs"
	"backend/utils/info"
)

func main() {
	info.Init()
	fs.Init()
	fiber.Init()
}
