package cmd

import (
	"log"
	"net/http"
)

func Execute() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	addr := ":8080"
	log.Println("Server started at port" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}