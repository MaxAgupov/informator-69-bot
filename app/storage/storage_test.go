package storage

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

)

func TestExistedStore(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestStoreUpload_*.tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := os.Remove(tmpfile.Name()); err != nil {
			log.Print(err)
		}
	}()

	var storageData = []byte(`{"123":{"id":123},"456":{"id":456},"789":{"id":789}}`)

	if _, err := tmpfile.Write(storageData); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	store := NewStore(tmpfile.Name())

	_, ok := store.Cache[123]
	if !ok {
		t.Error("There is no entry: 123")
	}
	_, ok = store.Cache[456]
	if !ok {
		t.Error("There is no entry: 456")
	}
	_, ok = store.Cache[789]
	if !ok {
		t.Error("There is no entry: 786")
	}
	_, ok = store.Cache[111]
	if ok {
		t.Error("There is extra entry: 111")
	}
}

func TestExtraFields(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestStoreUpload_*.tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := os.Remove(tmpfile.Name()); err != nil {
			log.Print(err)
		}
	}()

	var storageData = []byte(`{"123":{"id":123, "extra":1234},"456":{"id":456, "extra":4},"789":{"id":789, "extra":24}}`)

	if _, err := tmpfile.Write(storageData); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	store := NewStore(tmpfile.Name())

	_, ok := store.Cache[123]
	if !ok {
		t.Error("There is no entry: 123")
	}
	_, ok = store.Cache[456]
	if !ok {
		t.Error("There is no entry: 456")
	}
	_, ok = store.Cache[789]
	if !ok {
		t.Error("There is no entry: 786")
	}
	_, ok = store.Cache[111]
	if ok {
		t.Error("There is extra entry: 111")
	}
}