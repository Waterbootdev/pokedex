package main

import (
	"slices"
	"strings"
)

func CleanInput(text string) []string {
	return slices.DeleteFunc(strings.Split(strings.ToLower(text), " "), func(s string) bool {
		return s == ""
	})
}
