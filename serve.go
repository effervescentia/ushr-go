package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func locationSearch(w http.ResponseWriter, r *http.Request) {
	key := "AIzaSyCJx_xq8Sdv89KNRKrNWJFPdyMwcr5SCIk"
	uri := "https://maps.googleapis.com/maps/api/place/autocomplete/json"
	query := "a"

	callback := r.FormValue("jsoncallback")
	queryArg := r.FormValue("query")

	if queryArg != "" {
		query = queryArg
		log.Print("incoming query: " + query)
	}

	encodedURI := uri + "?input=" + url.QueryEscape(query) + "&types=(cities)&key=" + key
	log.Print("outgoing uri: " + encodedURI)
	response, err := http.Get(encodedURI)

	if err != nil {
		errorOut(w, err, callback)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			errorOut(w, err, callback)
		}
		print(w, callback, "%s\n", string(contents))
	}
}

func errorOut(w http.ResponseWriter, err error, callback string) {
	print(w, callback, "Error: %s", err)
	os.Exit(1)
}

func print(w http.ResponseWriter, callback string, format string, a ...interface{}) {
	template := format
	if callback != "" {
		w.Header().Set("Content-Type", "application/javascript")
		template = callback + "(" + format + ")"
	} else {
		w.Header().Set("Content-Type", "application/json")
	}

	fmt.Fprintf(w, template, a)
}

func main() {
	http.HandleFunc("/locations", locationSearch)
	http.ListenAndServe(":8000", nil)
}
