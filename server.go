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

func printNames(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()               // parse arguments, you have to call this by yourself
	fmt.Println("form", r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

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
	http.HandleFunc("/", printNames)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
