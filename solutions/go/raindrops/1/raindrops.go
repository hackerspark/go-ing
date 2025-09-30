package raindrops

import "fmt"

func Convert(number int) string {
    output := ""
    var isDivisibleBy3, isDivisibleBy5, isDivisibleBy7 bool
	if number % 3 == 0 {
        output += "Pling"
        isDivisibleBy3 = true
    }
    if number % 5 == 0 {
        output += "Plang"
        isDivisibleBy5 = true
    }
    if number % 7 == 0 {
        output += "Plong"
        isDivisibleBy7 = true
    }
    if !isDivisibleBy3 && !isDivisibleBy5 && !isDivisibleBy7 {
        output += fmt.Sprintf("%d", number)
    }
    return output
}
