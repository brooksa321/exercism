package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	//Convert dashes to spaces
	s = strings.Replace(s, "-", " ", -1)
	//Create empty string slice
	var firstLetters []string
	// Split string into a slice of words
	result := strings.SplitAfter(s, " ")

	//Grab the first letter of each word and capitalize it
	for _, words := range result {
		firstLetters = append(firstLetters, strings.ToUpper(string(words[0])))
	}

	//Create string from slice of first capitalized letters.
	abbreviation := strings.Join(firstLetters, "")

	return abbreviation
}
