package info

import (
	"backend/hub"
	"encoding/json"
	"os"
)

func Init() error {
	// * Unmarshal file
	var raw hub.Model
	file, _ := os.ReadFile("./data.json")
	err := json.Unmarshal(file, &raw)
	if err != nil {
		return err
	}

	// * assign hub
	hub.Hub = &raw
	println("Setup: data.json loaded")
	return nil
}
