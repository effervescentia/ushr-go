package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBasicSearch(t *testing.T) {
	w := DoSearch("")

	AssertContains(w.Body.String(), "\"predictions\"", t)
}

func TestQuery(t *testing.T) {
	w := DoSearch("query=ajax")

	AssertContains(w.Body.String(), "\"Ajax, ON, Canada\"", t)
}

func TestQueryWithSpace(t *testing.T) {
	w := DoSearch("query=new+york")

	AssertContains(w.Body.String(), "\"New York, NY, United States\"", t)
}

func AssertContains(actual string, expected string, t *testing.T) {
	if !strings.Contains(actual, expected) {
		t.Fail()
	}
}

func DoSearch(queryString string) *httptest.ResponseRecorder {
	handler := locationSearch

	req, err := http.NewRequest("GET", "localhost:8000/locations?"+queryString, nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	return w
}
