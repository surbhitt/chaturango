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

type Board *[8][8]Piece

var EMPTY_PIECE = Piece{name: "", color: ""}
