package natsio

import (
	"fmt"

	"github.com/thoas/go-funk"
)

// mergeAndUniqueArrayStrings ...
func mergeAndUniqueArrayStrings(arr1, arr2 []string) []string {
	var result = make([]string, 0)
	result = append(result, arr1...)
	result = append(result, arr2...)
	result = funk.UniqString(result)
	return result
}

// generateSubjectNames ...
func generateSubjectNames(streamName string, subjects []string) []string {
	var result = make([]string, 0)
	for _, subject := range subjects {
		name := combineStreamAndSubjectName(streamName, subject)
		result = append(result, name)
	}
	return result
}

func combineStreamAndSubjectName(stream, subject string) string {
	return fmt.Sprintf("%s.%s", stream, subject)
}
