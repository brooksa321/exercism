package twofer

import (
	"fmt"
)

// ShareWith should have a comment documenting it.
func ShareWith(s string) string {
	var name string
	if len(s) < 1 {
		name = "you"
	} else {
		name = s
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
