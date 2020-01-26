package utils

import "time"

import "strings"

// Line is map of header and data
type Line map[string]string

// String reads string value from the line with corresponding key.
func (l Line) String(k string) string {
	s, _ := l[k]
	return StringCleaner(s)
}

// StringArr reads multiple string value from the line with corresponding key.
func (l Line) StringArr(k string) []string {
	s, _ := l[k]
	if s == "" {
		return nil
	}
	return strings.Split(s, "  ")
}

// CleanStringArr reads multiple string value with cleanup from the line with corresponding key.
func (l Line) CleanStringArr(k string) []string {
	s, _ := l[k]
	ss := []string{}
	for _, x := range strings.Split(s, "  ") {
		if x == "" {
			continue
		}
		ss = append(ss, StringCleaner(x))
	}
	return ss
}

// Int reads integer value from the line with corresponding key.
func (l Line) Int(k string) int {
	s, _ := l[k]
	return IntCleaner(s)
}

// Uint reads uint value from the line with corresponding key.
func (l Line) Uint(k string) uint {
	return uint(l.Int(k))
}

// Time reads time value from the line with corresponding key.
func (l Line) Time(k string) time.Time {
	s, _ := l[k]
	return StringToTime(s)
}

// Bool reads boolean value from the line with corresponding key.
func (l Line) Bool(k string) bool {
	s, _ := l[k]
	if StringCleaner(s) == "yes" {
		return true
	}
	return false
}
