package storage

import (
	"fmt"
	kv "github.com/raymundovr/kvstore/core"
)

// One pointer to share among the package
var ServiceStorage *KVLogStorage

func InitializeEventsStorage() error {
	fmt.Println("[Storage] Initializing Events Storage");
	// We'll reuse this
	var err error

	logStorage, err := NewKVLogStorage("transactions.log")
	if err != nil {
		return fmt.Errorf("could not initialize transactions file: %w", err)
	}

	// Initialize global pointer
	ServiceStorage = logStorage

	events, errors := ServiceStorage.LoadEvents()
	// We'll reuse the same variables set
	event, isChannelOpen := KVStorageEvent{}, true

	fmt.Println("[Storage] Loading events from storage", isChannelOpen, err);
	for isChannelOpen && err == nil {
		select {
		// the <-channel syntax allows isChannelOpen to get 'false' when channel is closed
		// the consequent `case` here are not like those in a `switch`
		case err, isChannelOpen = <-errors:
		case event, isChannelOpen = <-events:
			switch event.EventType {
			case DeleteEvent:
				err = kv.Delete(event.Key)
			case PutEvent:
				err = kv.Put(event.Key, event.Value)
			}
		}
	}

	fmt.Println("[Storage] Running storage");
	ServiceStorage.Run()

	return err
}
