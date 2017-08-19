package main

import (
	"github.com/ProjectMOA/go2048/board"
	"github.com/ProjectMOA/go2048/input"
)

func main() {
	gameBoard := board.New(4, 4)

	for !gameBoard.IsFull() {
		gameBoard.Draw()
		direction, err := input.Get()
		if err != nil {
			break
		}
		gameBoard.Move(direction)
	}

}
