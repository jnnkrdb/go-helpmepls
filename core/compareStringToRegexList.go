package core

import (
	"regexp"
)

// compare a string with a list of regexp-like strings, if the
// list contains an expression, which matches with the given string, the
// return value will be true
//
// if an error occurs, the returnvalue will be false
//
// Parameters:
//   - `compare` : string > string, which should be compared to the list
//   - `listOfRegexp` : []string > list of regexp-like strings
func FindStringInRegexpList(compare string, listOfRegexp []string) (bool, int) {

	for index := range listOfRegexp {

		if match, err := regexp.MatchString(listOfRegexp[index], compare); err == nil {

			if match {

				return true, index
			}
		}
	}

	return false, -1
}
