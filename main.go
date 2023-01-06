package main

import (
	"fmt"
	"strings"

	"github.com/Tchayo/var-replace/libs/message"
	"github.com/Tchayo/var-replace/libs/template"
)

func main() {
	messages := []string{"Message for the ages", "Now you know"}
	// data structure the template will be applied to
	vars := make(map[string]interface{})
	vars["Name"] = "Brienne OBrien"
	vars["Number"] = "0700000000"
	vars["Message"] = strings.Join(messages, ", ")


	// process a template string
	resultA := template.ProcessString("Hello {{.Name}}, Your mobile number is {{.Number}}. {{.Message}}", vars)

	// process a template file
	// resultB := processFile("templates/got.tmpl", vars)

	// output
	fmt.Println(resultA)

	// test message to clean
	example := "#1 Sample Message with Special Characters (!@#$%^&*[])?!'><,|\\"

	// sanitize message using regexp

	processedString := message.SanitizeMessage(example)

	fmt.Println(example)
	fmt.Println(processedString)
	fmt.Printf("SMS count is %d \n", message.SMSCount(processedString))

}
