package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// StringCleaner converts string to lowercase.
func StringCleaner(s string) string {
	return strings.ToLower(s)
}

// IntCleaner converts string to int.
func IntCleaner(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// StringHeaderCleaner removes all the extra char from the string.
func StringHeaderCleaner(ss []string) []string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	for i, s := range ss {
		ss[i] = strings.ToLower(reg.ReplaceAllString(s, ""))
	}
	return ss
}

// StringToTime converts string to time.
func StringToTime(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}

// TypingSpeed parse long typing speed string to sort string.
func TypingSpeed(s string) string {
	switch s {
	case "40to60wordsminute":
		return "40<60"
	case "lessthan40wordsminute":
		return "<40"
	case "morethan60wordsminute":
		return "<60"
	}
	return ""
}
