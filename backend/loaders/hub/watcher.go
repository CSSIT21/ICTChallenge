package hub

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"

	"backend/utils/logger"
)

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
