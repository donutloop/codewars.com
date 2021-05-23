package Simple_Fun__1__Seats_in_Theater

func seatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - col + 1) * (nRows - row)
}