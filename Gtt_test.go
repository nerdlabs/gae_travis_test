package Gtt

import (
	"testing"

	"appengine/aetest"
	"appengine/datastore"
)

type Something struct {
	Foo int64
	Bar string
}

func TestSomethingInMyDatastore(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	sth := Something{Foo: 12345, Bar: "baz"}
	if _, err = datastore.Put(c, datastore.NewIncompleteKey(c, "Something", nil), &sth); err != nil {
		t.Fatal(err)
	}
}
