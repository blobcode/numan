package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	dir := os.Args[1]
	port := os.Args[2]
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	log.Println("Listening on http://localhost:" + port)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
