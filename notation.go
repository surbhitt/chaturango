package main

import "strconv"

// func get_piece() Piece {
// }

func parse_notation(notation string) []Coord {
	from_y := int(notation[0]) - 97
	from_x, _ := strconv.Atoi(string(notation[1]))
	to_y := int(notation[3]) - 97
	to_x, _ := strconv.Atoi(string(notation[4]))
	return []Coord{{from_x - 1, from_y}, {to_x - 1, to_y}}
}
