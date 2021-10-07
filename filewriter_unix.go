//go:build darwin || linux || netbsd || openbsd
// +build darwin linux netbsd openbsd

package files

import "strings"

var invalidChars = `/` + "\x00"

func isValidFilename(filename string) bool {
	return !strings.ContainsAny(filename, invalidChars)
}
