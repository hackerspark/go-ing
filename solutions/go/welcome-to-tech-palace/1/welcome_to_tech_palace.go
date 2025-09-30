package techpalace

import (
    "strings"
    "fmt"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    template := "%s\n%s\n%s"
    repeated := strings.Repeat("*", numStarsPerLine)
    value := fmt.Sprintf(template, repeated, welcomeMsg, repeated)
    fmt.Println(value)
	return value
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	afterRemoval := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(afterRemoval)
}
