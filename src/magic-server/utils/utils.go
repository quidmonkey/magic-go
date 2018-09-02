package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

// Check checks an error to see if it's defined, and if so, throws it
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// FileHandler reads a json file and sends a response
func FileHandler(filename string) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, err := ioutil.ReadFile(filename)
	Check(err)

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(file))
	}
}

// LoadFileHandlerRoutes loads the json files in a given directory and then
// registers a file handler for each of them
func LoadFileHandlerRoutes(router *httprouter.Router, jsonPath string) {
	fp := filepath.Join(jsonPath, "*.json")
	files, err := filepath.Glob(fp)
	Check(err)

	for _, file := range files {
		_, filename := filepath.Split(file)
		route := "/" + strings.Replace(filename, ".json", "", -1)
		router.GET(route, FileHandler(file))
		fmt.Println("Registered", route, "GET route")
	}
}

// RegisterHandler registers a handler with httprouter
// router 	httprouter object
// method 	HTTP method i.e. "GET", "POST", etc
// endpoint HTTP endpoint i.e. "/", etc
// handler 	Request handler
func RegisterHandler(
	router *httprouter.Router,
	method string,
	endpoint string,
	handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params),
) {
	h := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		uuid := uuid.NewV4()

		fmt.Println("~~~ Request #", uuid, "- Received", endpoint, "Request")
		handler(w, r, ps)
		fmt.Println("~~~ Request #", uuid, "- Completed", endpoint, "Request -", time.Since(start), "Elapsed")
	}

	router.Handle(method, endpoint, h)
}
