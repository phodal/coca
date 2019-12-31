package jpackage

import "strings"

func GetClassName(path string) string {
	split := strings.Split(path, ".")
	return strings.Join(split[:len(split)-1], ".")
}

func GetMethodName(path string) string {
	split := strings.Split(path, ".")
	return strings.Join(split[len(split)-1:], ".")
}

