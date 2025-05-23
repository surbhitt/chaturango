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

func friendly_fire(pos Coord, move Coord, board *Board) bool {
	attacker := board[pos.x][pos.y]
	attacked := board[move.x][move.y]
	// check if the sq attacked contains a friendly piece
	if attacked != EMPTY_PIECE {
		if attacked.color == attacker.color {
			return true
		}
	}

	// check for all other pieces except knight
	// if there is a friendly piece in line of attack
	if attacker.name != "knight" {
		dx := move.x - pos.x
		dy := move.y - pos.y
		if dx != 0 {
			if dx < 0 {
				dx = -1
			} else {
				dx = 1
			}
		}
		if dy != 0 {
			if dy < 0 {
				dy = -1
			} else {
				dy = 1
			}
		}

		cur_x := pos.x
		cur_y := pos.y
		end_x := move.x
		end_y := move.y

		for cur_x != end_x || cur_y != end_y {
			collateral := board[cur_x][cur_y]
			if collateral.position != attacker.position && collateral.color == attacker.color {
				return true
			}
			cur_x += dx
			cur_y += dy
		}
	}

	return false
}

func get_king_moves(pos Coord, board *Board) []Coord {
	var king_moves []Coord
	// create a square around the pos
	// all coords in that square except for the original pos
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) {
				move := Coord{pos.x + i, pos.y + j}
				if check_board_bound(move) && !friendly_fire(pos, move, board) {
					king_moves = append(king_moves, move)
				}
			}
		}
	}
	return king_moves
}

func get_queen_moves(pos Coord, board *Board) []Coord {
	var queen_moves []Coord
	// extrapolate diagonally, vertically and horizontally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y + i}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			queen_moves = append(queen_moves, move)
		}
		move = Coord{pos.x + i, pos.y - i}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			queen_moves = append(queen_moves, move)
		}
	}
	return queen_moves
}

func get_rook_moves(pos Coord, board *Board) []Coord {
	var rook_moves []Coord
	// extrapolate vertically and horizontally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + 0, pos.y + i}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			rook_moves = append(rook_moves, move)
		}
		move = Coord{pos.x + i, pos.y + 0}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			rook_moves = append(rook_moves, move)
		}
	}
	return rook_moves
}

func get_bishop_moves(pos Coord, board *Board) []Coord {
	var bishop_moves []Coord
	// extrapolate diagonally
	for i := -8; i < 8; i++ {
		move := Coord{pos.x + i, pos.y + i}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			bishop_moves = append(bishop_moves, move)
		}
		move = Coord{pos.x - i, pos.y + i}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			bishop_moves = append(bishop_moves, move)
		}
	}
	return bishop_moves
}

func get_knight_moves(pos Coord, board *Board) []Coord {
	var knight_moves []Coord
	possible_moves := [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {-1, 2}, {1, 2}, {-1, -2}, {1, -2}}
	for _, possible_move := range possible_moves {
		move := Coord{pos.x + possible_move[0], pos.y + possible_move[1]}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			knight_moves = append(knight_moves, move)
		}
	}
	return knight_moves
}

func get_pawn_moves(p Piece, board *Board) []Coord {
	// TODO: cross attack move
	pos := p.position
	var dif int
	if p.color == "black" {
		dif = -1
	} else {
		dif = 1
	}
	var moves []Coord
	if !p.has_moved {
		// first move 2 steps
		move := Coord{pos.x + 2*dif, pos.y}
		if check_board_bound(move) && !friendly_fire(pos, move, board) {
			moves = append(moves, move)
		}
	}
	move := Coord{pos.x + dif, pos.y}
	if check_board_bound(move) && !friendly_fire(pos, move, board) {
		moves = append(moves, move)
	}

	return moves
}

func (p Piece) get_valid_moves(board *Board) []Coord {
	switch p.name {
	case "king":
		return get_king_moves(p.position, board)
	case "queen":
		return get_queen_moves(p.position, board)
	case "rook":
		return get_rook_moves(p.position, board)
	case "bishop":
		return get_bishop_moves(p.position, board)
	case "knight":
		return get_knight_moves(p.position, board)
	case "pawn":
		// needs more context than just position
		return get_pawn_moves(p, board)
	default:
		return []Coord{}
	}
}

func add_pawns(board *Board, color string) {
	rank := 1
	if color == "black" {
		rank = 6
	}
	for file := range 8 {
		board[rank][file] = make_piece(color, "pawn", Coord{rank, file})
	}
}

func add_pieces(board *Board, color string) {
	rank := 0
	if color == "black" {
		rank = 7
	}
	board[rank][0] = make_piece(color, "rook", Coord{rank, 0})
	board[rank][7] = make_piece(color, "rook", Coord{rank, 7})
	board[rank][1] = make_piece(color, "knight", Coord{rank, 1})
	board[rank][6] = make_piece(color, "knight", Coord{rank, 6})
	board[rank][2] = make_piece(color, "bishop", Coord{rank, 2})
	board[rank][5] = make_piece(color, "bishop", Coord{rank, 5})
	board[rank][3] = make_piece(color, "queen", Coord{rank, 3})
	board[rank][4] = make_piece(color, "king", Coord{rank, 4})
}

func make_move(board *Board, from Coord, to Coord) {
	// check if the from and to are board bound
	if !check_board_bound(from) {
		fmt.Printf("%s %d [ERR] OUT OF BOARD\n", string('a'+from.y), from.x+1)
		return
	}
	if !check_board_bound(to) {
		fmt.Printf("%s %d [ERR] OUT OF BOARD\n", string('a'+to.y), to.x+1)
		return
	}

	// check if the from sq has a piece
	piece := board[from.x][from.y]
	if piece == EMPTY_PIECE {
		fmt.Printf("no piece found on the square %s %d\n", string('a'+from.y), from.x+1)
		return
	}

	// check if the move is valid for the piece on sq from
	valid_moves := piece.get_valid_moves(board)
	is_valid := false
	for _, valid_move := range valid_moves {
		if to.x == valid_move.x && to.y == valid_move.y {
			is_valid = true
			break
		}
	}

	if !is_valid {
		fmt.Println("Not a valid move for piece")
		return
	}

	piece.position = to
	piece.has_moved = true
	board[to.x][to.y] = piece
	// TODO: better way to represent emtpy piece
	board[from.x][from.y] = EMPTY_PIECE
}

func initiate_board() Board {
	var board Board
	for i := range 8 {
		for j := range 8 {
			board[i][j] = EMPTY_PIECE
		}
	}
	add_pawns(&board, "black")
	add_pawns(&board, "white")
	add_pieces(&board, "black")
	add_pieces(&board, "white")
	return board
}
