func isValidSudoku(board [][]byte) bool {

	hashMap := make(map[string]bool)

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {

			currValue := string(board[row][column])

			if currValue == "." {
				continue
			}
			_, rowOk := hashMap[currValue+"found in row"+string(row)]
			_, columnOk := hashMap[currValue+"found in column"+string(column)]
			_, gridOk := hashMap[currValue+"found in grid"+string(row/3)+"-"+string(column/3)]

			if rowOk || columnOk || gridOk {
				return false
			} else {
				hashMap[currValue+"found in row"+string(row)] = true
				hashMap[currValue+"found in column"+string(column)] = true
				hashMap[currValue+"found in grid"+string(row/3)+"-"+string(column/3)] = true
			}
		}
	}
	return true
}