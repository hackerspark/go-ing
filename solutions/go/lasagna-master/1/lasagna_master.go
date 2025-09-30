package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2
    }
    return len(layers) * avgPrepTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodlesQuantity int, sauceQuantity float64) {
    for _, val := range layers {
        if val == "sauce" {
            sauceQuantity += 0.2
        } else if val == "noodles" {
            noodlesQuantity += 50
        }
    }
    return 
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
    actualScale := float64(scale) / 2

    var scaledQuantities []float64

    for _, val := range quantities {
       scaledQuantities = append(scaledQuantities, val * actualScale)
    }

    return scaledQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
