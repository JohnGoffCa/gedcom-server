package main

import (
	"bytes"
	"fmt"
	"github.com/iand/gedcom"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func printGedcom(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("testdata/sample.ged")

	d := gedcom.NewDecoder(bytes.NewReader(data))

	g, _ := d.Decode()

	for _, rec := range g.Individual {
		if len(rec.Name) > 0 {
			fmt.Fprintln(w, rec.Name[0].Name)
		}
	}
}

func main() {
	http.HandleFunc("/", printGedcom)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
