package Gtt

import (
	"log"
	"testing"
	"time"

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
	time.Sleep(1 * time.Second)
	q := datastore.NewQuery("Something")
	var sths []Something
	_, err = q.GetAll(c, &sths)
	if len(sths) != 1 {
		t.Errorf("len(somethings): expected: 1, actual: %d\n", len(sths))
	}
}

func TestThatThisNotWillFailAndBlockThePullRequest(t *testing.T) {
	log.Printf("I love to fail!^Wsucceeeeeed!\n")
}
