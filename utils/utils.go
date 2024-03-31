package utils

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func WithoutNils[T any](slice []*T) []T {
	var result []T
	for _, item := range slice {
		if item != nil {
			result = append(result, *item)
		}
	}
	return result
}

// PrettyPrintError prints an error in a pretty way
// It splits on `details: ` and adds `due to: \n`
func PrettyPrintError(err error) {
	all := make([]string, 0)

	currentErr := err
	for errors.Unwrap(currentErr) != nil {
		all = append(all, currentErr.Error())
		currentErr = errors.Unwrap(currentErr)
	}

	all = append(all, currentErr.Error())

	for i := 0; i < len(all); i++ {
		// remove the parts of the string that exists in all[i+1]
		if i+1 < len(all) {
			all[i] = strings.ReplaceAll(all[i], all[i+1], "")
		}
	}

	for i := 0; i < len(all); i++ {
		// remove details: from the string
		all[i] = strings.ReplaceAll(all[i], "details: ", "")

		// remove punctuation in the end of the string
		all[i] = strings.TrimRight(all[i], ".")
	}

	output := ""

	for i := 0; i < len(all); i++ {
		if i == 0 {
			output = all[i]
		} else {
			output = fmt.Sprintf("%s\n\tdue to: %s", output, all[i])
		}
	}

	log.Println(output)
}
