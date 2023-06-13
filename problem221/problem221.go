package problem221

func maximalSquare(matrix [][]byte) int {
	return memoization(matrix)
}

func memoization(matrix [][]byte) int {
	rowNumber, columnNumber := len(matrix), len(matrix[0])
	memo := make([][]int, rowNumber)

	for row := 0; row < rowNumber; row++ {
		memo[row] = make([]int, columnNumber)
		for column := 0; column < columnNumber; column++ {
			memo[row][column] = -1
		}
	}

	var dp func(row, column int) int
	dp = func(row, column int) int {
		if row >= rowNumber || column >= columnNumber || matrix[row][column] == 0 {
			return 0
		}

		if memo[row][column] != -1 {
			return memo[row][column]
		}

		right, bottom, diagonal := dp(row, column+1), dp(row+1, column), dp(row+1, column+1)

		memo[row][column] = 1 + min(min(right, bottom), diagonal)

		return memo[row][column]
	}

	maxSquare := 0

	for row := 0; row < rowNumber; row++ {
		for column := 0; column < columnNumber; column++ {
			result := dp(row, column)
			maxSquare = max(maxSquare, result*result)
		}
	}

	return maxSquare
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
