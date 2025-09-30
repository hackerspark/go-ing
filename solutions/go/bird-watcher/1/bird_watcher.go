package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    totalCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
        totalCount += birdsPerDay[i]
    }
    return totalCount
}

func BirdsInWeek(birdsPerDay []int, week int) int {
    start := 7 * (week - 1)
    end := start + 6
    totalCount := 0

    // Limit end to last index of birdsPerDay slice
    if end >= len(birdsPerDay) {
        end = len(birdsPerDay) - 1
    }

    for i := start; i <= end; i++ {
        totalCount += birdsPerDay[i]
    }
    return totalCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i< len(birdsPerDay); i++ {
        if i % 2 == 0 {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
