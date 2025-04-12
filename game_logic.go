package main

import "fmt"

// using inverted colors because the black pieces look more white
func get_char_repr(piece_name string, color string) rune {
	switch piece_name {
	case "king":
		if color == "white" {
			return '♚' // \u265A
		} else {
			return '♔' // \u2654
		}
	case "queen":
		if color == "white" {
			return '♛' // \u265B
		} else {
			return '♕' // \u2655
		}
	case "rook":
		if color == "white" {
			return '♜' // \u265C
		} else {
			return '♖' // \u2656
		}
	case "bishop":
		if color == "white" {
			return '♝' // \u265D
		} else {
			return '♗' // \u2657
		}
	case "knight":
		if color == "white" {
			return '♞' // \u265E
		} else {
			return '♘' // \u2658
		}
	case "pawn":
		if color == "white" {
			return '♟' // \u265F
		} else {
			return '♙' // \u2659
		}
	default:
		return '#'
	}
}

func make_piece(color string, name string, position Coord) Piece {
	repr := get_char_repr(name, color)
	has_moved := false
	return Piece{
		color:     color,
		name:      name,
		position:  position,
		has_moved: has_moved,
		repr:      repr,
	}
}

func check_board_bound(pos Coord) bool {
	return pos.x < 8 && pos.y < 8 && pos.x > -1 && pos.y > -1
}

func get_king_moves(pos Coord) []Coord {
	var king_moves []Coord
	// create a square around the pos
	// all coords in that square except for the original pos
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) {
				move := Coord{pos.x + i, pos.y + j}
				if check_board_bound(move) {
					king_moves = append(king_moves, move)
				}
			}
		}
	}
	return king_moves
}

func get_queen_moves(pos Coord) []Coord {
	var queen_moves []Coord
	// extrapolate diagonally, vertically and horizontally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y + i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y - i}
		if check_board_bound(move) {
			queen_moves = append(queen_moves, move)
		}
	}
	return queen_moves
}

func get_rook_moves(pos Coord) []Coord {
	var rook_moves []Coord
	// extrapolate vertically and horizontally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) {
			rook_moves = append(rook_moves, move)
		}
		move = Coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) {
			rook_moves = append(rook_moves, move)
		}
	}
	return rook_moves
}

func get_bishop_moves(pos Coord) []Coord {
	var bishop_moves []Coord
	// extrapolate diagonally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + i, pos.y + i}
		if check_board_bound(move) {
			bishop_moves = append(bishop_moves, move)
		}
		move = Coord{pos.x - i, pos.y + i}
		if check_board_bound(move) {
			bishop_moves = append(bishop_moves, move)
		}
	}
	return bishop_moves
}

func get_knight_moves(pos Coord) []Coord {
	var knight_moves []Coord
	possible_moves := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {-1, 2}, {1, 2}, {-1, -2}, {1, -2}}
	for i := 0; i < len(possible_moves); i++ {
		move := Coord{pos.x + possible_moves[i][0], pos.y + possible_moves[i][1]}
		if check_board_bound(move) {
			knight_moves = append(knight_moves, move)
		}
	}
	return knight_moves
}

func get_pawn_moves(p Piece) []Coord {
	// TODO: check for first move and allow 2 steps
	// TODO: cross attack move
	pos := p.position
	var dif int
	if p.color == "black" {
		dif = -1
	} else {
		dif = 1
	}
	var moves []Coord
	if p.has_moved != false {
		// first move 2 steps
		move := Coord{pos.x + 2*dif, pos.y}
		if check_board_bound(move) {
			moves = append(moves, move)
		}
	}
	move := Coord{pos.x + dif, pos.y}
	if check_board_bound(move) {
		moves = append(moves, move)
	}

	return moves
}

func (p Piece) get_valid_moves() []Coord {
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
		// needs more context than just position
		return get_pawn_moves(p)
	default:
		return []Coord{}
	}
}

func add_pawns(board *[8][8]Piece, color string) {
	if color == "black" {
		for i := 0; i < 8; i++ {
			board[6][i] = make_piece(color, "pawn", Coord{6, i})
		}
	}
	if color == "white" {
		for i := 0; i < 8; i++ {
			board[1][i] = make_piece(color, "pawn", Coord{1, i})
		}
	}
}

func add_pieces(board *[8][8]Piece, color string) {
	if color == "black" {
		board[7][0] = make_piece(color, "rook", Coord{0, 0})
		board[7][7] = make_piece(color, "rook", Coord{0, 7})
		board[7][1] = make_piece(color, "knight", Coord{0, 1})
		board[7][6] = make_piece(color, "knight", Coord{0, 6})
		board[7][2] = make_piece(color, "bishop", Coord{0, 2})
		board[7][5] = make_piece(color, "bishop", Coord{0, 5})
		board[7][3] = make_piece(color, "queen", Coord{0, 3})
		board[7][4] = make_piece(color, "king", Coord{0, 4})
	}
	if color == "white" {
		board[0][0] = make_piece(color, "rook", Coord{7, 0})
		board[0][7] = make_piece(color, "rook", Coord{7, 7})
		board[0][1] = make_piece(color, "knight", Coord{7, 1})
		board[0][6] = make_piece(color, "knight", Coord{7, 6})
		board[0][2] = make_piece(color, "bishop", Coord{7, 2})
		board[0][5] = make_piece(color, "bishop", Coord{7, 5})
		board[0][3] = make_piece(color, "queen", Coord{7, 3})
		board[0][4] = make_piece(color, "king", Coord{7, 4})
	}
}

func make_move(board *[8][8]Piece, from Coord, to Coord) {
    fmt.Println(from.x, from.y)
    fmt.Println(to.x, to.y)
	piece := board[from.x][from.y]
	piece.position = to
	board[to.x][to.y] = piece
    // TODO: better way to represent emtpy piece
	board[from.x][from.y] = EMPTY_PIECE
}

func initiate_board() [8][8]Piece {
	var board [8][8]Piece
	add_pawns(&board, "black")
	add_pawns(&board, "white")
	add_pieces(&board, "black")
	add_pieces(&board, "white")
	return board
}
