package main

import "fmt"

func parse_notation(notation string) []Coord {
	// given notation in the form
	// a1pa3 convert to []Coord{from, to}
	return []Coord{
		a1not_to_xycoord(notation[:2]),
		a1not_to_xycoord(notation[3:]),
	}
}

func a1not_to_xycoord(notation string) Coord {
	// convert a1 -> Coord{0,0}
	// TODO guard against values out of bound
	if len(notation) != 2 {
		panic("invalid notation: must be 2 characters like 'a1'")
	}
	return Coord{
		x: int(notation[1] - '1'),
		y: int(notation[0] - 'a'),
	}
}

func xycoord_to_a1not(coord Coord) string {
	// convert Coord{0,0} -> a1
	// TODO guard against values out of bound
	file := string('a' + coord.y)
	rank := fmt.Sprintf("%d", coord.x+1)
	return file + rank
}
