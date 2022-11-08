package hub

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"

	"backend/types/extend"
	"backend/utils/logger"
)

func Init() {
	Hub = &Model{
		Teams:  nil,
		Topics: nil,
		AdminConn: &extend.ConnModel{
			Context: "ADMIN_CONN",
			Conn:    nil,
			Mutex:   &sync.Mutex{},
		},
		LeaderboardProjectorConn: &extend.ConnModel{
			Context: "LEADERBOARD_PROJECTOR_CONN",
			Conn:    nil,
			Mutex:   &sync.Mutex{},
		},
		CardProjectorConn: &extend.ConnModel{
			Context: "CARD_PROJECTOR_CONN",
			Conn:    nil,
			Mutex:   &sync.Mutex{},
		},
	}
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
	Hub.Topics = raw.Topics
	Hub.Teams = raw.Teams

	logger.Log(logrus.Debug, "LOADED HUB DATA")

	return nil
}

func Watch() {
	// * Create new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Log(logrus.Panic, "UNABLE TO START WATCHER: "+err.Error())
	}

	// * Start listening for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {
					if err := Load(); err != nil {
						logger.Log(logrus.Warn, "UNABLE TO LOAD HUB: "+err.Error())
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Log(logrus.Warn, "WATCHER ERROR: "+err.Error())
			}
		}
	}()

	// * Add a watch path
	err = watcher.Add("./data.json")
	if err != nil {
		log.Fatal(err)
	}
}
