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

/** fixName
	* Properly formats names to remove slashes and escape quotation marks
 **/
func fixName(name string) string {
	fixed := strings.Replace(name, "/", "", -1)
	fixed = strings.Replace(fixed, `"`, `\"`, -1)
	return fixed
}

/** familyList
	* writes JSON representing the entire list of families from GEDCOM
 **/
func familyList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "[")
	for i, rec := range g.Family {
		if i != 0 {
			fmt.Fprint(w, ",")
		}
		fmt.Fprintf(w, `"%s"`, rec.Xref)
	}
	fmt.Fprintln(w, "]")
}

func family(w http.ResponseWriter, r *http.Request) {
	for _, rec := range g.Family {
		if id := vestigo.Param(r, "id"); !strings.EqualFold(id, rec.Xref) {
			continue
		}

		fmt.Fprintln(w, "{")
		fmt.Fprintf(w, `"id":"%s",`, rec.Xref)
		fmt.Fprintf(w, `"Husband":"%s",`, fixName(rec.Husband.Xref))
		fmt.Fprintf(w, `"Wife":"%s",`, fixName(rec.Wife.Xref))
		fmt.Fprintln(w, `"Children":[`)
		for i, child := range rec.Child {
			if i != 0 {
				fmt.Fprint(w, ",")
			}
			if len(child.Name) > 0 {
				fmt.Fprintf(w, `"%s"`, fixName(child.Xref))
			}
		}
		fmt.Fprintln(w, "]")
		fmt.Fprintln(w, "}")
	}
}

/** individualList
	* writes JSON representing the entire list of individuals from GEDCOM
 **/
func individualList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "[")
	for i, rec := range g.Individual {
		if i != 0 {
			fmt.Fprint(w, ",")
		}
		fmt.Fprintf(w, `"%s`, rec.Xref)
	}
	fmt.Fprintln(w, "]")
}

/** individual
	* writes JSON containing the information about an individual
	* to the client.
 **/
func individual(w http.ResponseWriter, r *http.Request) {
	for _, rec := range g.Individual {
		if id := vestigo.Param(r, "id"); !strings.EqualFold(id, rec.Xref) {
			continue
		}

		fmt.Fprint(w, "{")
		fmt.Fprintf(w, `"id":"%s",`, rec.Xref)
		fmt.Fprintf(w, `"name":"%s",`, fixName(rec.Name[0].Name))
		fmt.Fprintf(w, `"sex":"%s",`, rec.Sex)
		fmt.Fprint(w, `"events":[`)
		for i, evnt := range rec.Event {
			if i != 0 {
				fmt.Fprint(w, ",")
			}
			fmt.Fprint(w, "{")
			fmt.Fprintf(w, `"event":"%s",`, evnt.Tag)
			fmt.Fprintf(w, `"date":"%s",`, evnt.Date)
			fmt.Fprintf(w, `"place":"%s"`, evnt.Place.Name)
			fmt.Fprint(w, "}")
		}
		fmt.Fprint(w, `],`)
		fmt.Fprint(w, `"attributes":[`)
		for i, attr := range rec.Attribute {
			if i != 0 {
				fmt.Fprint(w, ",")
			}
			fmt.Fprint(w, "{")
			fmt.Fprintf(w, `"tag":"%s",`, attr.Tag)
			fmt.Fprintf(w, `"value":"%s"`, attr.Value)
			fmt.Fprint(w, "}")
		}
		fmt.Fprint(w, `],`)
		fmt.Fprint(w, `"mother":"`)
		if len(rec.Parents) > 0 {
			fmt.Fprint(w, fixName(rec.Parents[0].Family.Wife.Xref))
		}
		fmt.Fprint(w, `",`)
		fmt.Fprint(w, `"father":"`)
		if len(rec.Parents) > 0 {
			fmt.Fprint(w, fixName(rec.Parents[0].Family.Husband.Xref))
		}
		fmt.Fprint(w, `",`)
		fmt.Fprint(w, `"children":[`)
		if len(rec.Family) > 0 {
			for i, child := range rec.Family[0].Family.Child {
				if len(child.Name) > 0 {
					if i != 0 {
						fmt.Fprint(w, ",")
					}
					fmt.Fprintf(w, `"%s"`, fixName(child.Xref))
				}
			}
		}
		fmt.Fprint(w, "]}")
	}
}
