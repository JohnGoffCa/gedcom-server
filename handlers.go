package main

import (
	"bytes"
	"fmt"
	"github.com/husobee/vestigo"
	"github.com/iand/gedcom"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var g = parseGedcomFile()

func parseGedcomFile() *gedcom.Gedcom {
	data, err := ioutil.ReadFile("assets/gedcom/sample.ged")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	d := gedcom.NewDecoder(bytes.NewReader(data))

	g, err := d.Decode()
	if err != nil {
		log.Fatal("Error decoding GEDCOM:", err)
	}

	return g
}

func fixName(name string) string {
	fixed := strings.Replace(name, "/", "", -1)
	fixed = strings.Replace(fixed, "\"", "\\\"", -1)
	return fixed
}

func familyList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "{")
	for i, rec := range g.Family {
		if i != 0 {
			fmt.Fprint(w, ",")
		}
		fmt.Fprintf(w, `"%s":["%s","%s"]`,
			rec.Xref,
			fixName(rec.Husband.Name[0].Name),
			fixName(rec.Wife.Name[0].Name),
		)
	}
	fmt.Fprintln(w, "}")
}

func family(w http.ResponseWriter, r *http.Request) {
	for _, rec := range g.Family {
		if id := vestigo.Param(r, "id"); !strings.EqualFold(id, rec.Xref) {
			continue
		}

		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, `"id":"`+rec.Xref+`",`)
		fmt.Fprintln(w, `"Husband":"`+fixName(rec.Husband.Name[0].Name)+`",`)
		fmt.Fprintln(w, `"Wife":"`+fixName(rec.Wife.Name[0].Name)+`",`)
		fmt.Fprintln(w, `"Children":[`)
		for i, child := range rec.Child {
			if i != 0 {
				fmt.Fprint(w, ",")
			}
			if len(child.Name) > 0 {
				fmt.Fprintln(w, `"`+fixName(child.Name[0].Name)+`"`)
			}
		}
		fmt.Fprintln(w, "]")
		fmt.Fprintln(w, "}")
	}
}

func individualList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "{")
	for i, rec := range g.Individual {
		if i != 0 {
			fmt.Fprint(w, ",")
		}
		fmt.Fprintf(w, `"%s":"%s"`, rec.Xref, fixName(rec.Name[0].Name))
	}
	fmt.Fprintln(w, "}")
}

func individual(w http.ResponseWriter, r *http.Request) {
	for _, rec := range g.Individual {
		if id := vestigo.Param(r, "id"); !strings.EqualFold(id, rec.Xref) {
			continue
		}

		fmt.Fprintln(w, rec.Xref)
		fmt.Fprintln(w, fixName(rec.Name[0].Name))
		fmt.Fprintln(w, rec.Sex)
		for _, evnt := range rec.Event {
			fmt.Fprintln(w, evnt.Tag)
			fmt.Fprintln(w, evnt.Date)
			fmt.Fprintln(w, evnt.Place.Name)
		}
		for _, attr := range rec.Attribute {
			fmt.Fprintln(w, attr.Tag)
			fmt.Fprintln(w, attr.Value)
		}
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Parents: ")
		fmt.Fprintln(w, fixName(rec.Family[0].Family.Husband.Name[0].Name))
		fmt.Fprintln(w, fixName(rec.Family[0].Family.Wife.Name[0].Name))
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Children:")
		for _, child := range rec.Family[0].Family.Child {
			if len(child.Name) > 0 {
				fmt.Fprintln(w, fixName(child.Name[0].Name))
			}
		}
		fmt.Fprintln(w, "")
	}
}
