package main

import (
	"strings"
)

var SERIAL = "SERIAL"
var CONCURRENT = "CONCURRENT"

func isValidType(downloadType string) bool {

	if strings.Compare(downloadType, SERIAL) == 0 {
		return true
	}
	if strings.Compare(downloadType, CONCURRENT) == 0 {
		return true
	}
	return false
}

func isURLsEmpty(URLs []string) bool {

	return len(URLs) == 0
}
