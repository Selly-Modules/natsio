package natsio

import "github.com/thoas/go-funk"

// mergeAndUniqueArrayStrings ...
func mergeAndUniqueArrayStrings(arr1, arr2 []string) []string {
	var result = make([]string, 0)
	result = append(result, arr1...)
	result = append(result, arr2...)
	result = funk.UniqString(result)
	return result
}
