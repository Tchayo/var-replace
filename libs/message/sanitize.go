package message

import (
	"log"
	"regexp"
)

// SanitizeMessage : remove special characters and whatnots
func SanitizeMessage(s string) string {
	var processedString string
	// sanitize message using regexp
	reg, err := regexp.Compile("[^A-Za-z0-9_*?!@#%*()-/\\s]")
	if err != nil {
		log.Fatal(err)
	}
	processedString = reg.ReplaceAllString(s, "")
	return processedString
}

// SMSCount : return message count. Max chars for single message 160
func SMSCount(s string) int {
	var processedCount int

	processedCount = 1
	wordCount := len(s)

	if wordCount > 160 {
		processedCount = int(wordCount / 160)
	}
	return processedCount
}
