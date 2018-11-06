package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tupkung/gophercises/exercise_2/urlshort"
)

var (
	port     = flag.String("port", ":9999", "Port for listening")
	yamlFile = flag.String("yamlfile", "./config-url.yml", "Path of YAML file")
)

func main() {
	flag.Parse()
	mux := defaultMux()
	yamlToUrls, err := ioutil.ReadFile(*yamlFile)
	if err != nil {
		log.Fatal(err)
	}
	yamlHandler, err := urlshort.YAMLHandler(yamlToUrls, mux)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening at port %s", *port)
	log.Fatal(http.ListenAndServe(*port, yamlHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
