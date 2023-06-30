package main

import (
	"fmt"

	"github.com/raymundovr/kvstore/storage"
)

// One pointer to share among the package
var ServiceStorage *storage.KVLogStorage

func initializeEventsStorage() error {
	fmt.Println("[Storage] Initializing Events Storage");
	// We'll reuse this
	var err error

	logStorage, err := storage.NewKVLogStorage("transactions.log")
	if err != nil {
		return fmt.Errorf("could not initialize transactions file: %w", err)
	}

	// Initialize global pointer
	ServiceStorage = logStorage

	events, errors := ServiceStorage.LoadEvents()
	// We'll reuse the same variables set
	event, isChannelOpen := storage.KVStorageEvent{}, true

	fmt.Println("[Storage] Loading events from storage");
	for isChannelOpen && err != nil {
		select {
		// the <-channel syntax allows isChannelOpen to get 'false' when channel is closed
		// the consequent `case` here are not like those in a `switch`
		case err, isChannelOpen = <-errors:
		case event, isChannelOpen = <-events:
			switch event.EventType {
			case storage.DeleteEvent:
				err = Delete(event.Key)
			case storage.PutEvent:
				err = Put(event.Key, event.Value)
			}
		}
	}

	fmt.Println("[Storage] Running storage");
	ServiceStorage.Run()

	return err
}
