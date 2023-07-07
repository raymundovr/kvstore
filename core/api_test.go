package core

import (
	"errors"
	"testing"
)

type TestStorage struct {}

func (s *TestStorage) WriteDelete(k string) error {
	return nil
}
func (s *TestStorage) WritePut(k, v string) error {
	return nil
}

func (s *TestStorage) Run() {

}
func (s *TestStorage) LoadEvents() (<-chan KVStorageEvent, <-chan error) {
	evchan := make(chan KVStorageEvent)
	errchan := make(chan error)

	return evchan, errchan
}

func (s *TestStorage) Err() <-chan error {
	errchan := make(chan error)

	return errchan
}

func TestPut(t *testing.T) {
	s := NewKVStore(&TestStorage{})
	err := s.Put("test", "kvstore")

	if err != nil {
		t.Fatal("Cannot PUT key")
	}
}

func TestGet(t *testing.T) {
	s := NewKVStore(&TestStorage{})
	putKeyValue(t, "test", "kvstore", s)

	v, err := s.Get("test")
	if err != nil {
		t.Fatal("Cannot GET", err)
	}

	if v != "kvstore" {
		t.Fatal("GET incorrect value returned")
	}

	v, err = s.Get("undefined")
	if !errors.Is(err, ErrorNoSuchKey) || v != "" {
		t.Fatal("GET incorrect NoSuchKey assertion", err, v)
	}
}

func TestDelete(t *testing.T) {
	s := NewKVStore(&TestStorage{})
	putKeyValue(t, "test", "deleteme", s)

	err := s.Delete("test")
	if err != nil {
		t.Fatal("Cannot Delete")
	}
}

func putKeyValue(t *testing.T, k, v string, s *KVStore) {
	err := s.Put(k, v)

	if err != nil {
		t.Fatal("Cannot PUT key for test")
	}
}