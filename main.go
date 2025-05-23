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
	print_board(board)
	var notation string
	// var moves []Move
	for {
		fmt.Scan(&notation)
		if notation == "m" {
			var coord string
			fmt.Scan(&coord)
			print_moves(board, coord)
			continue
		} else {
			from_to := parse_notation(notation)
			make_move(&board, from_to[0], from_to[1])
			print_board(board)
		}
	}
	// TODO
	// you click on a piece and get the valid moves list
	// you make a valid move
}
