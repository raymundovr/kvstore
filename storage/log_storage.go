package storage

import (
	"bufio"
	"fmt"
	"os"
	"github.com/raymundovr/kvstore/core"
)

const LINE_FORMAT = "%d\t%d\t%s\t%s\n"

type KVLogStorage struct {
	events       chan<- core.KVStorageEvent // Write-only (to the channel)
	errors       <-chan error          //Read-only (from the channel)
	lastSequence uint64
	file         *os.File
}

func NewKVLogStorage(filename string) (core.KVStorage, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open log file: %w", err)
	}

	return &KVLogStorage{file: file}, nil
}

func (l *KVLogStorage) WritePut(k, v string) error {
	l.events <- core.KVStorageEvent{EventType: core.PutEvent, Key: k, Value: v}
	return nil
}

func (l *KVLogStorage) WriteDelete(k string) error {
	l.events <- core.KVStorageEvent{EventType: core.DeleteEvent, Key: k}
	return nil
}

func (l *KVLogStorage) Err() <-chan error {
	return l.errors
}

func (l *KVLogStorage) Run() {
	events := make(chan core.KVStorageEvent, 16)
	errors := make(chan error, 1)

	l.events = events
	l.errors = errors

	go func() {
		for ev := range events {
			l.lastSequence++

			_, err := fmt.Fprintf(
				l.file,
				LINE_FORMAT,
				l.lastSequence, ev.EventType, ev.Key, ev.Value)

			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

func (l *KVLogStorage) LoadEvents() (<-chan core.KVStorageEvent, <-chan error) {
	scanner := bufio.NewScanner(l.file)
	// We declare a channel of concrete/copied values, not pointers
	events := make(chan core.KVStorageEvent)
	errors := make(chan error)

	go func() {
		// We reuse the same event
		var event core.KVStorageEvent

		defer close(events)
		defer close(errors)

		for scanner.Scan() {
			line := scanner.Text()

			if _, err := fmt.Sscanf(line, LINE_FORMAT, &event.Sequence, &event.EventType, &event.Key, &event.Value); err != nil {
				errors <- fmt.Errorf("error while reading from log: %w", err)
				return
			}

			if l.lastSequence >= event.Sequence {
				errors <- fmt.Errorf("transaction numbers out of order")
				return
			}

			l.lastSequence = event.Sequence
			// Hence, we can send the same event variable down the line
			events <- event
		}

		if err := scanner.Err(); err != nil {
			errors <- fmt.Errorf("error reading log file: %w", err)
			return
		}
	}()

	return events, errors
}

func (logStorage *KVLogStorage) Load(store *core.KVStore) error {
	fmt.Println("[Storage] Initializing Events Storage");
	// We'll reuse this
	var err error

	events, errors := logStorage.LoadEvents()
	// We'll reuse the same variables set
	event, isChannelOpen := core.KVStorageEvent{}, true

	fmt.Println("[Storage] Loading events from storage", isChannelOpen, err);
	for isChannelOpen && err == nil {
		select {
		// the <-channel syntax allows isChannelOpen to get 'false' when channel is closed
		// the consequent `case` here are not like those in a `switch`
		case err, isChannelOpen = <-errors:
		case event, isChannelOpen = <-events:
			switch event.EventType {
			case core.DeleteEvent:
				err = store.Delete(event.Key)
			case core.PutEvent:
				err = store.Put(event.Key, event.Value)
			}
		}
	}

	fmt.Println("[Storage] Running storage");
	ServiceStorage.Run()

	return err
}
