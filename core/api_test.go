package core

import (
	"errors"
	"testing"
)

func TestPut(t *testing.T) {
	err := Put("test", "kvstore")

	if err != nil {
		t.Fatal("Cannot PUT key")
	}
}

func TestGet(t *testing.T) {
	putKeyValue(t, "test", "kvstore")

	v, err := Get("test")
	if err != nil {
		t.Fatal("Cannot GET", err)
	}

	if v != "kvstore" {
		t.Fatal("GET incorrect value returned")
	}

	v, err = Get("undefined")
	if !errors.Is(err, ErrorNoSuchKey) || v != "" {
		t.Fatal("GET incorrect NoSuchKey assertion", err, v)
	}
}

func TestDelete(t *testing.T) {
	putKeyValue(t, "test", "deleteme")

	err := Delete("test")
	if err != nil {
		t.Fatal("Cannot Delete")
	}
}

func putKeyValue(t *testing.T, k, v string) {
	err := Put(k, v)

	if err != nil {
		t.Fatal("Cannot PUT key for test")
	}
}