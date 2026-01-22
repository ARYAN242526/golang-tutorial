package greet

import "strings"

// exported functions start with capital letter
// other packages can call it  -> greet.Hello()

func Hello(name string) string {
	clean := normalizeName(name)

	return "Hello , " + clean
}

func normalizeName(name string) string {

	n := strings.TrimSpace(name)

	if n == "" {
		return "Guest"
	}
	return strings.ToUpper(n)
}