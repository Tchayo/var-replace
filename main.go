package main

import (
	"fmt"

	"github.com/Tchayo/var-replace/libs/message"
	"github.com/Tchayo/var-replace/libs/template"
)

func main() {
	// data structure the template will be applied to
	vars := make(map[string]interface{})
	vars["Name"] = "Brienne OBrien"
	vars["Number"] = "0700000000"
	vars["Message"] = []string{"Message for the ages", "Now you know"}

	// process a template string
	resultA := template.ProcessString("Hello {{.Name}}, My message to you is {{.Number}} has been compromized. {{.Extra}}", vars)

	// process a template file
	// resultB := processFile("templates/got.tmpl", vars)

	// output
	fmt.Println(resultA)

	// test message to clean
	example := "#1 message to all f@$ck$r out here! is [loud] to stop(decist)..say what?!'><,|\\"

	// sanitize message using regexp

	processedString := message.SanitizeMessage(example)

	fmt.Println(example)
	fmt.Println(processedString)
	fmt.Printf("SMS count is %d \n", message.SMSCount(processedString))

}
