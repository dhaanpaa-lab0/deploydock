package utilities

import (
	"os"
	"strings"
)

func GetEnvWithDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// InvertStringMap takes a map with string keys and string values and returns a new map with the keys and values inverted.
func InvertStringMap(m map[string]string) map[string]string {
	invertedMap := make(map[string]string)
	for key, value := range m {
		invertedMap[value] = key
	}
	return invertedMap
}

// FilterMapByKeyPrefix filters a map based on whether a string key starts with a given prefix, case-insensitively.
// Returns a new map containing only the keys that match the condition.
func FilterMapByKeyPrefix(inputMap map[string]string, prefix string) map[string]string {
	filteredMap := make(map[string]string)

	for key, value := range inputMap {
		if strings.HasPrefix(strings.ToLower(key), strings.ToLower(prefix)) {
			filteredMap[key] = value
		}
	}

	return filteredMap
}

func UniqueStrings(strings []string) []string {
	uniqueMap := make(map[string]bool) // Map to track unique strings
	var result []string                // Slice to hold unique strings

	for _, str := range strings {
		if _, exists := uniqueMap[str]; !exists {
			uniqueMap[str] = true        // Mark string as seen
			result = append(result, str) // Add unique string to result slice
		}
	}

	return result
}

func ParsePath(path string) string {
	// Check if tilde character is present
	if path[0] == '~' {
		// Get the home directory
		home := os.Getenv("HOME")
		// Replace the tilde with the home directory
		path = home + path[1:]
	} else if path[0] == '.' {
		// Get the current directory
		wd, _ := os.Getwd()
		// Replace the dot with the current directory
		path = wd + path[1:]
	} else {
		// Do nothing
		return path
	}
	return path
}
