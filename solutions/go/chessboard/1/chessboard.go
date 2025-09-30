package chessboard

// File represents a single column on a chessboard.
type File []bool

// Chessboard represents a chessboard with 8 rows and 8 columns.
type Chessboard map[string]File

// CountInFile returns the number of occupied squares in a file.
func CountInRank(cb Chessboard, file int) int {
	count := 0
	for _, rank := range cb {
		if file >= 1 && file <= len(rank) && rank[file-1] {
			count++
		}
	}
	return count
}

// CountInRank returns the number of occupied squares in a rank.
func CountInFile(cb Chessboard, rank string) int {
	count := 0
	if file, ok := cb[rank]; ok {
		for _, square := range file {
			if square {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return 64
}

// CountOccupied returns the number of occupied squares on the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for rank := range cb {
		count += CountInFile(cb, rank)
	}
	return count
}