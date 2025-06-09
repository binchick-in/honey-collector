package utils

import "strings"

// PreparePorts parses a comma-separated list of ports and returns a slice of trimmed port strings.
func PreparePorts(x string) []string {
	ports := strings.Split(x, ",")
	var a []string
	for _, i := range ports {
		a = append(a, strings.TrimSpace(i))
	}
	return a
}
