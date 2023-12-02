package stringutils

import "strings"

func TrimAndSplit(input string, sep string) []string {
	trimmed := strings.TrimSpace(input)
	return strings.Split(trimmed, sep)
}
