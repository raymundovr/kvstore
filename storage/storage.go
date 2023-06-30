package storage

type KVStorage interface {
	WriteDelete(k string) error
	WritePut(k, v string) error

	Run()
	LoadEvents() (chan KVStorageEvent, chan error)

	Err() <-chan error
}

type KVStorageEvent struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

type EventType byte

const (
	_                  = iota
	PutEvent EventType = iota
	DeleteEvent
)
