package main

import (
	"fmt"
)

type coord struct {
	x, y int
}

type piece struct {
	color     string
	name      string
	position  coord
	has_moved bool
	repr      rune
}

func get_char_repr(piece_name string) rune {
	switch piece_name {
	case "king":
		return 'k'
	case "queen":
		return 'q'
	case "rook":
		return 'r'
	case "bishop":
		return 'b'
	case "knight":
		return 'n'
	case "pawn":
		return 'p'
	default:
		return '#'
	}
}

func make_piece(color string, name string, position coord) piece {
	repr := get_char_repr(name)
	has_moved := false
	return piece{
		color:     color,
		name:      name,
		position:  position,
		has_moved: has_moved,
		repr:      repr,
	}
}

// TODO: board bouond checks
func check_board_bound(pos coord) bool {
	return pos.x < 8 && pos.y < 8 && pos.x > -1 && pos.y > -1
}

func get_king_moves(pos coord) []coord {
	var king_moves []coord
	// create a square around the pos
	// all coords in that square except for the original pos
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) {
				move := coord{pos.x + i, pos.y + j}
				if check_board_bound(move) {
					king_moves = append(king_moves, move)
				}
			}
		}
	}
	return king_moves
}

func get_queen_moves(pos coord) []coord {
	var queen_moves []coord
	// extrapolate diagonally, vertically and horizontally
	for i := -8; i < 8; i++ {
		move := coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = coord{pos.x + i, pos.y + i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = coord{pos.x + i, pos.y - i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
	}
	return queen_moves
}

func get_rook_moves(pos coord) []coord {
	var rook_moves []coord
	// extrapolate vertically and horizontally
	for i := -8; i < 8; i++ {
		move := coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) {
			rook_moves = append(rook_moves, move)
		}
		move = coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) {
			rook_moves = append(rook_moves, move)
		}
	}
	return rook_moves
}

func get_bishop_moves(pos coord) []coord {
	var bishop_moves []coord
	// extrapolate diagonally
	for i := -8; i < 8; i++ {
		move := coord{pos.x + i, pos.y + i}
		if check_board_bound(move) {
			bishop_moves = append(bishop_moves, move)
		}
		move = coord{pos.x - i, pos.y + i}
		if check_board_bound(move) {
			bishop_moves = append(bishop_moves, move)
		}
	}
	return bishop_moves
}

func get_knight_moves(pos coord) []coord {
	var knight_moves []coord
	possible_moves := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {-1, 2}, {1, 2}, {-1, -2}, {1, -2}}
	for i := 0; i < len(possible_moves); i++ {
		move := coord{pos.x + possible_moves[i][0], pos.y + possible_moves[i][1]}
		if check_board_bound(move) {
			knight_moves = append(knight_moves, move)
		}
	}
	return knight_moves
}

func get_pawn_moves(pos coord) []coord {
	// TODO: check for first move and allow 2 steps
	// TODO: cross attack move
	move := coord{pos.x-1, pos.y}
	if check_board_bound(move) {
		return []coord{move}
	} else {
		return []coord{move}
	}
}

func (p piece) get_valid_moves() []coord {
	switch p.name {
	case "king":
		return get_king_moves(p.position)
	case "queen":
		return get_queen_moves(p.position)
	case "rook":
		return get_rook_moves(p.position)
	case "bishop":
		return get_bishop_moves(p.position)
	case "knight":
		return get_knight_moves(p.position)
	case "pawn":
		return get_pawn_moves(p.position)
	default:
		return []coord{}
	}
}

// func initiate_board() {}
func print_board(p piece) {
	valid_moves := p.get_valid_moves()

	var board [8][8]rune

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = '-'
		}
	}

	for coord_i := 0; coord_i < len(valid_moves); coord_i++ {
		x := valid_moves[coord_i].x
		y := valid_moves[coord_i].y
		board[x][y] = '0'
	}

	board[p.position.x][p.position.y] = p.repr

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// board := initiate_board()
	p := make_piece("white", "knight", coord{x: 4, y: 4})
	print_board(p)
}
