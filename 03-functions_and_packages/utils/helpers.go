package utils

import "strings"

// ToUpper converts a string to uppercase.
// This function is exported because its name starts with a capital letter.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// toLower converts a string to lowercase.
// This function is unexported because its name starts with a lowercase letter.
// It can only be used by other functions within the 'utils' package.
func toLower(s string) string {
	return strings.ToLower(s)
}
