package logs

import (
    "strings"
    "unicode/utf8"
    )

var Lookup = map[rune]string{
    '❗': "recommendation",
    '🔍': "search",
    '☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
        if _, ok := Lookup[c]; ok {
            return Lookup[c]
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var sb strings.Builder

    for _, val := range log {
        if val == oldRune {
            sb.WriteRune(newRune)
        } else {
            sb.WriteRune(val)
        }
    }

    return sb.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
