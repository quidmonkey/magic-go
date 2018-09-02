package main

import (
	"fmt"
	"log"
	"magic-server/utils"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// HealthHandler serves up a health endpoint that determins if the server is running
func HealthHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Ok\n")
}

func main() {
	var port string

	if len(os.Args) == 2 {
		port = os.Args[1]
	} else {
		port = "8000"
	}

	router := httprouter.New()
	server := "localhost:" + port

	utils.RegisterHandler(router, "GET", "/health", HealthHandler)
	fmt.Println("Registered /health GET route")

	utils.LoadFileHandlerRoutes(router, "../../data")

	fmt.Println("~~~ Running Server on Port", port)
	log.Fatal(http.ListenAndServe(server, router))
}
