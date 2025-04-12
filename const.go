package main

type Coord struct {
	x, y int
}

type Piece struct {
	color     string
	name      string
	position  Coord
	has_moved bool
	repr      rune
}

type Move struct {
	piece       Piece
	destination Coord
}

var EMPTY_PIECE = Piece{name: "", color: ""}
