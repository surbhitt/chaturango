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
	var command string
	for {
		fmt.Scan(&command)
		if command == "isready" {
			/// ok
		} else if command == "validmoves" {
			var notation string
			fmt.Scan(&notation)
			coord := a1not_to_xycoord(notation)
			piece := board[coord.x][coord.y]
			moves := piece.get_valid_moves(&board)
			for _, move := range moves {
				fmt.Println(xycoord_to_a1not(move))
			}
		} else if command == "makemove" {
			var from, to string
			fmt.Scan(&from)
			fmt.Scan(&to)
			from_coord := a1not_to_xycoord(from)
			to_coord := a1not_to_xycoord(to)
			fmt.Println(from)
			fmt.Println(to)
			make_move(&board, from_coord, to_coord)
			print_board(board)
			// TODO should indicate if the move is invalid
		} else {
			fmt.Println("[WARN] Unindentified command", command)
		}
		fmt.Println("ok")
	}
}
