package isogram

import (
    "slices"
    "strings"
    )

func IsIsogram(word string) bool {
    var lookup []rune
    word = strings.ToLower(word)
	for _, ch := range word {
        if ch == '-' || ch == ' ' {
            continue
        }
        if slices.Contains(lookup, ch) {
            return false
        } else {
            lookup = append(lookup, ch)
        }
    }
    return true
}
