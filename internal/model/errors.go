package model

import "strings"

type ValidationErrors []error

// for logs
func (v ValidationErrors) Error() string {
	var sb strings.Builder
	for i, err := range v {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}

func (v ValidationErrors) Errors() []string {
	result := make([]string, len(v))

	for i, err := range v {
		result[i] = err.Error()
	}
	return result
}
