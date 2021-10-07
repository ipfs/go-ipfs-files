//go:build windows
// +build windows

package files

import "strings"

var invalidChars = `<>:"/\|?*` + "\x00"

var reservedNames = map[string]struct{}{
	"CON":  {},
	"PRN":  {},
	"AUX":  {},
	"NUL":  {},
	"COM1": {},
	"COM2": {},
	"COM3": {},
	"COM4": {},
	"COM5": {},
	"COM6": {},
	"COM7": {},
	"COM8": {},
	"COM9": {},
	"LPT1": {},
	"LPT2": {},
	"LPT3": {},
	"LPT4": {},
	"LPT5": {},
	"LPT6": {},
	"LPT7": {},
	"LPT8": {},
	"LPT9": {},
}

func isValidFilename(filename string) bool {
	_, isReservedName := reservedNames[filename]
	return !strings.ContainsAny(filename, invalidChars) &&
		!isReservedName
}
