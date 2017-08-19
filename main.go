package main

import (
	"fmt"

	"github.com/ProjectMOA/go2048/board"
	"github.com/ProjectMOA/go2048/input"
)

func main() {
	gameBoard := board.New(4, 4)

	for !gameBoard.IsFull() {
		gameBoard.Draw()
		direction, err := input.Get()
		if err != nil {
			fmt.Print(err)
			break
		}
		gameBoard.Move(direction)
	}

}
