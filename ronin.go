package main

import (
	"github.com/howeyc/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				log.Println("event:", event)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("./test")
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}
