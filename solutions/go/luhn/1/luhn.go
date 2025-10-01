package luhn

import (
    "unicode"
    "strconv"
    "strings"
)

func Valid(id string) bool {
    // Remove spaces from input
    id = strings.ReplaceAll(id, " ", "")
    if len(id) <= 1 {
        return false
    }

    var input []int
    // Convert runes to integers
    for _, ch := range id {
        if !unicode.IsDigit(ch) {
            return false
        }
        val, _ := strconv.Atoi(string(ch))
        input = append(input, val)
    }

    total := 0
    n := len(input)
    // Start doubling every second digit from the right
    for i := n - 1; i >= 0; i-- {
        val := input[i]
        if (n-1-i)%2 == 1 {
            val *= 2
            if val > 9 {
                val -= 9
            }
        }
        total += val
    }

    return total%10 == 0
}
