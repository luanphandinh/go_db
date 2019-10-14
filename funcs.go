package dbs

import (
	"fmt"
	"regexp"
)

func concatStrings(values []string, glue string) (s string) {
	if glue == "" {
		for _, value := range values {
			s += value
		}

		return s
	}

	for index, value := range values {
		if value == "" {
			continue
		}

		if index == 0 {
			s += value
		} else {
			s += glue + value
		}
	}

	return s
}

func inStringArray(needle string, values []string) bool  {
	for _, value := range values {
		if value == needle {
			return true
		}
	}

	return false
}

// fmt.Sprintf("%#v", arg) will return a Go-syntax representation of the value
// eg: []string{"1", "2"}
// This func will get content inside {} and return as string
func getContentOutOfArraySyntax(arg interface{}) string {
	splitter := regexp.MustCompile("[{}]")
	parsedArg := splitter.Split(fmt.Sprintf("%#v", arg), -1)
	re := regexp.MustCompile(`"`)
	return re.ReplaceAllString(parsedArg[1], `'`)
}
