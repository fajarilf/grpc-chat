package utils

import "strings"

func SearchQuery(slice []string, query string) []string {
	results := []string{}
	query = strings.ToLower(query)
	query = query[1:]

	for _, item := range slice {
		if strings.Contains(strings.ToLower(item), query) {
			results = append(results, item)
		}
	}

	return results
}
