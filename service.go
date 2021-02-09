package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/projectorangejuice/jnotes/frontend"
)

var (
	port int
)

func main() {

	flag.IntVar(&port, "port", 9000, "Port for server to run on")
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/", frontend.IndexHandler)
	//router.HandleFunc("/post", api.PlayerHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
	}
}
