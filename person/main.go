package main

import (
	"net/http"
	"reflect"

	"github.com/go-msvc/crud"
	"github.com/go-msvc/errors"
	person "github.com/go-msvc/ms-person"
	"github.com/go-msvc/store/mongo"
)

func main() {
	mc := mongo.Config{Database: "ms-person"}
	personStore, err := mc.New("person", reflect.TypeOf(person.Person{}))
	if err != nil {
		panic(errors.Wrapf(err, "failed to create person store"))
	}
	nationalityStore, err := mc.New("nationality", reflect.TypeOf(person.Nationality{}))
	if err != nil {
		panic(errors.Wrapf(err, "failed to create nationality store"))
	}

	crud := crud.New().With(personStore).With(nationalityStore)
	crud.AddToMux(http.DefaultServeMux)
	http.ListenAndServe("localhost:8000", http.DefaultServeMux)
}
