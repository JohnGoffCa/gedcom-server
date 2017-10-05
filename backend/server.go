package main

import (
	"fmt"
	"github.com/husobee/vestigo"
	"log"
	"net/http"
	"os"
)

func main() {
	ch := make(chan byte, 1)
	router := vestigo.NewRouter()
	var port string
	if p := os.Getenv("PORT"); p != "" {
		port = p
	} else {
		port = "9090"
	}

	////////////
	// ROUTES //
	////////////
	router.Get("/*", http.FileServer(http.Dir("./assets")).ServeHTTP)
	router.Get("/api/individual", individualList)
	router.Get("/api/individual/:id", individual)
	router.Get("/api/family", familyList)
	router.Get("/api/family/:id", family)

	//run server in goroutine to print after the server has started
	go func() {
		log.Fatal("ListenAndServe:", http.ListenAndServe(":"+port, router))
		ch <- 1
	}()

	fmt.Println("Listening on port :" + port)
	<-ch //prevent main function from exiting automatically
}
