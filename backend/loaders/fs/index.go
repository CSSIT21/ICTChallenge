package fs

import (
	"backend/types/database"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

func Init() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)

					// * Unmarshal file
					raw := new(database.Raw)
					file, _ := os.ReadFile("./data.json")
					err = json.Unmarshal(file, &raw)
					if err != nil {
						log.Fatal(err)
					}
					spew.Dump(raw)

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("./data.json")
	if err != nil {
		log.Fatal(err)
	}
}
