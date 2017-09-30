package main

import (
	"bytes"
	"fmt"
	"github.com/husobee/vestigo"
	"github.com/iand/gedcom"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func printGedcom(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("assets/gedcom/sample.ged")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	d := gedcom.NewDecoder(bytes.NewReader(data))

	g, err := d.Decode()
	if err != nil {
		log.Fatal("Error decoding GEDCOM:", err)
	}

	for _, rec := range g.Family {
		fmt.Fprintln(w, rec.Xref)
		fmt.Fprintln(w, "Husband: "+rec.Husband.Name[0].Name)
		fmt.Fprintln(w, "Wife: "+rec.Wife.Name[0].Name)
		fmt.Fprintln(w, "Children:")
		for _, child := range rec.Child {
			if len(child.Name) > 0 {
				fmt.Fprintln(w, child.Name[0].Name)
			}
		}
		fmt.Fprintln(w, "")
	}
}

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
	router.Get("/", http.FileServer(http.Dir("./assets")).ServeHTTP)
	router.Get("/api/individual/:id", printGedcom)
	router.Get("/api/family/:id", printGedcom)

	//run server in goroutine to print after the server has started
	go func() {
		log.Fatal("ListenAndServe:", http.ListenAndServe(":"+port, router))
		ch <- 1
	}()

	fmt.Println("Listening on port :" + port)
	<-ch //prevent main function from exiting automatically
}
