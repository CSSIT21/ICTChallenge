package hub

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"

	"backend/utils/logger"
)

func Init() {
	if err := Load(); err != nil {
		logger.Log(logrus.Panic, "UNABLE TO LOAD HUB: "+err.Error())
	}
	Watch()
}

func Load() error {
	// * Unmarshal file
	var raw *Model
	file, _ := os.ReadFile("./data.json")
	err := json.Unmarshal(file, &raw)
	if err != nil {
		return err
	}

	// * Assign hub
	Hub = raw

	logger.Log(logrus.Debug, "LOADED HUB DATA")

	return nil
}
