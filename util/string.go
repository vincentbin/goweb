package util

import "strings"

func ParseUrl(url string) []string {
	vs := strings.Split(url, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item == "" {
			continue
		}
		parts = append(parts, item)
	}
	return parts
}
