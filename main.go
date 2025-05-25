package main

import (
	"fmt"
	"strconv"
)

func print_board(board Board) {
	var board_rep [8][8]rune

	for i := range 8 {
		for j := range 8 {
			board_rep[i][j] = '-'
		}
	}
	for i := range 8 {
		for j := range 8 {
			if (board[i][j] != Piece{}) {
				board_rep[i][j] = board[i][j].repr
			}
		}
	}
	for i := range 8 {
		for j := range 8 {
			fmt.Printf("%c ", board_rep[i][j])
		}
		fmt.Println()
	}
}

func print_moves(board Board, coord string) {
	var board_rep [8][8]rune

	for i := range 8 {
		for j := range 8 {
			board_rep[i][j] = '-'
		}
	}
	for i := range 8 {
		for j := range 8 {
			if (board[i][j] != Piece{}) {
				board_rep[i][j] = board[i][j].repr
			}
		}
	}
	x := int(coord[0] - 'a')
	y, err := strconv.Atoi(string(coord[1]))
	if err == nil {
		y--
	}
	piece := board[x][y]
	moves := piece.get_valid_moves(&board)
	for _, move := range moves {
		board_rep[move.x][move.y] = '#'
	}

	for i := range 8 {
		for j := range 8 {
			fmt.Printf("%c ", board_rep[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := initiate_board()
	var notation string
	for {
		fmt.Scan(&notation)
		coord := a1not_to_xycoord(notation)
		piece := board[coord.x][coord.y]
		moves := piece.get_valid_moves(&board)
		for _, move := range moves {
			fmt.Println(xycoord_to_a1not(move))
		}
		fmt.Println("readyok")
	}
}

// func main() {
// 	test := Coord{7, 7}
// 	fmt.Println(xycoord_to_a1not(test))
// 	fmt.Println(a1not_to_xycoord(xycoord_to_a1not(test)).x, a1not_to_xycoord(xycoord_to_a1not(test)).y)
// }
