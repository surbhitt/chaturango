package main

import (
	"fmt"
)

func print_board(board [8][8]Piece) {
	var board_rep [8][8]rune

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board_rep[i][j] = '-'
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (board[i][j] != Piece{}) {
				board_rep[i][j] = board[i][j].repr
			}
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board_rep[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := initiate_board()
	print_board(board)
    var notation string
    // var moves []Move
    for {
        fmt.Scan(&notation)
        from_to := parse_notation(notation)
        make_move(&board, from_to[0], from_to[1])
	    print_board(board)
    }
    // TODO
    // you click on a piece and get the valid moves list
    // you make a valid move
}
