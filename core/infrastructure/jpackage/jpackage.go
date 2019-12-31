package jpackage

import "strings"

func GetClassName(child string) string {
	split := strings.Split(child, ".")
	return strings.Join(split[:len(split)-1], ".")
}

func GetMethodName(child string) string {
	split := strings.Split(child, ".")
	return strings.Join(split[len(split)-1:], ".")
}

