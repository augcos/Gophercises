package main

import (
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	
	"github.com/augcos/Gophercises/URLShortener/handler"
)

func main() {
	// we define and parse the flag for the YAML or JSON file with the url data
	dataFilename := flag.String("filename", "url.yaml", "file with the data in YAML or JSON format (path, url)")
	flag.Parse()

	// we create a mux as the default fallback
	mux := defaultMux()
	// we build a default http handler that works with a map object
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)	
	
	// we read the data from the file in the flag
	data, err := ioutil.ReadFile(*dataFilename)
	if err != nil {
		panic(err)
	}

	// we check if the file has a .yaml or .json extension
	if strings.Contains(*dataFilename,".yaml") {
		// we use the data from the yaml file to create the http handler
		dataHandler, err := handler.YAMLHandler([]byte(data), mapHandler)
		if err != nil {
			panic(err)
		}	
		// we start the http server on port 8080
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", dataHandler)

	} else if strings.Contains(*dataFilename,".json") {
		// we use the data from the json file to create the http handler
		dataHandler, err := handler.JSONHandler([]byte(data), mapHandler)
		if err != nil {
			panic(err)
		}	
		// we start the http server on port 8080
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", dataHandler)
	}
}

// defaultMux function: creates a http mux to act as handler
func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

// hello function: this returns the hello world string with for any request
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}