package day4

// gets randomized boards and numbers for the winners
func bingo(file string) {
	// first line in file are the winning numbers
	// everything else = boards
	input := convertDataToArray(file)

}

/*
boards := {1: [[22, 13, 17, 11, 0], [...], ...], ...}
bitwiseBoards := {1: [[0, 0, 0, 0, 0], [...], ...]}
row = 10000

isWinner = 11111
if row & isWinner == isWinner => 10000 == 11111 which is false
*can store entire board in 32bit int

winningBoard = 11111111111111111111111111111
bitwise to shift every len(row) from row[0] to get the first bit of the next row

column mask + row mask are needed

mask = bit sequence of winning condition (i.e. condition to check against)

when a value is matched on boards, set that index on bitwise to 1
*/
