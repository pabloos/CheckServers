package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	serverDependencies := Servers{
		{
			Name: "MongoDB Server",
			URL:  "mongo",
			Port: 27017,
		},
		{
			Name: "MySQL Server",
			URL:  "mysql",
			Port: 3306,
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})

	CheckServers(serverDependencies)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
