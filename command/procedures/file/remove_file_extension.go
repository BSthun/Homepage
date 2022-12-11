package file

import "strings"

func RemoveFileExtension(name string) string {
	// Find last dot
	dot := strings.LastIndex(name, ".")

	// If dot is not found, return original string
	if dot == -1 {
		return name
	}

	// Return string without extension
	return name[:dot]
}
