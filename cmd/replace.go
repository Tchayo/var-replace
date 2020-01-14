package main

import (
	"fmt"

	"github.com/tchayo/var-replace/libs/template"
)

func main() {
	// data structure the template will be applied to
	vars := make(map[string]interface{})
	vars["Name"] = "Brienne OBrien"
	vars["Number"] = "0725836628"
	vars["Message"] = []string{"Message for the ages", "Now you know"}

	// process a template string
	resultA := template.processString("Hello {{.Name}}, My message to you is {{.Number}} has been compromized. {{.Extra}}", vars)

	// process a template file
	// resultB := processFile("templates/got.tmpl", vars)

	// output
	fmt.Println(resultA)

}
