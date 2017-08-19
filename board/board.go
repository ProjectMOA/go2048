package board

import (
	"math/rand"

	tm "github.com/buger/goterm"
)

// Board is the game board
type Board struct {
	Width, Height int8
	Values        [][]int32
}

// Direction represents a direction
type Direction int

// Define the different directions
const (
	Up Direction = iota
	Down
	Left
	Right
)

// New creates a new game board
func New(width, height int8) Board {
	rand.Seed(42) // Initialize random number generation
	b := Board{Width: width, Height: height}
	// Initialize the slices
	b.Values = make([][]int32, height, height)
	for i := int8(0); i < height; i++ {
		b.Values[i] = make([]int32, width, width)
	}
	b.addRandomValue()
	return b
}

// Draw draws the board on the terminal, erasing the previous board
func (b *Board) Draw() {
	// Clear the terminal
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println()
	tm.Println()
	tm.Println()
	tm.Println()
	for i := int8(0); i < b.Height; i++ {
		for j := int8(0); j < b.Width; j++ {
			tm.Printf("|%d", b.Values[i][j])
		}
		tm.Println("|")
	}

	tm.Flush() // Call it every time at the end of rendering
}

// IsFull returns true if the board is full
func (b *Board) IsFull() bool {
	for i := int8(0); i < b.Height; i++ {
		for j := int8(0); j < b.Width; j++ {
			if b.Values[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// Move moves the values in the board in the selected direction
func (b *Board) Move(dir Direction) {
	changes := true
	switch dir {
	case Up:
		for changes {
			changes = false
			for col := int8(0); col < b.Width; col++ {
				for row := int8(1); row < b.Height; row++ {
					if b.Values[row-1][col] == 0 && b.Values[row][col] != 0 {
						// Exchange values
						b.Values[row-1][col], b.Values[row][col] = b.Values[row][col], b.Values[row-1][col]
						changes = true
					}
					if b.Values[row-1][col] != 0 && b.Values[row-1][col] == b.Values[row][col] {
						// Combine values
						b.Values[row-1][col], b.Values[row][col] = b.Values[row][col]*2, 0
						changes = true
					}
				}
			}
		}
	case Down:
		for changes {
			changes = false
			for col := int8(0); col < b.Width; col++ {
				for row := int8(b.Height - 2); row >= 0; row-- {
					if b.Values[row+1][col] == 0 && b.Values[row][col] != 0 {
						// Exchange values
						b.Values[row+1][col], b.Values[row][col] = b.Values[row][col], b.Values[row+1][col]
						changes = true
					}
					if b.Values[row+1][col] != 0 && b.Values[row+1][col] == b.Values[row][col] {
						// Combine values
						b.Values[row+1][col], b.Values[row][col] = b.Values[row][col]*2, 0
						changes = true
					}
				}
			}
		}
	case Left:
		for changes {
			changes = false
			for row := int8(0); row < b.Height; row++ {
				for col := int8(1); col < b.Width; col++ {
					if b.Values[row][col-1] == 0 && b.Values[row][col] != 0 {
						// Exchange values
						b.Values[row][col-1], b.Values[row][col] = b.Values[row][col], b.Values[row][col-1]
						changes = true
					}
					if b.Values[row][col-1] != 0 && b.Values[row][col-1] == b.Values[row][col] {
						// Combine values
						b.Values[row][col-1], b.Values[row][col] = b.Values[row][col]*2, 0
						changes = true
					}
				}
			}
		}
	case Right:
		for changes {
			changes = false
			for row := int8(0); row < b.Height; row++ {
				for col := int8(b.Width - 2); col >= 0; col-- {
					if b.Values[row][col+1] == 0 && b.Values[row][col] != 0 {
						// Exchange values
						b.Values[row][col+1], b.Values[row][col] = b.Values[row][col], b.Values[row][col+1]
						changes = true
					}
					if b.Values[row][col+1] != 0 && b.Values[row][col+1] == b.Values[row][col] {
						// Combine values
						b.Values[row][col+1], b.Values[row][col] = b.Values[row][col]*2, 0
						changes = true
					}
				}
			}
		}
	}
	b.addRandomValue()
}

func (b *Board) addRandomValue() {
	// Find empty spaces in the board
	tmpslices := make([][2]int8, 0, b.Height*b.Width)
	for i := int8(0); i < b.Height; i++ {
		for j := int8(0); j < b.Width; j++ {
			if b.Values[i][j] == 0 {
				tmpslices = append(tmpslices, [2]int8{i, j})
			}
		}
	}
	spaces := len(tmpslices)
	if spaces == 0 {
		return
	}
	// Put a 2 in a random space
	id := rand.Intn(len(tmpslices))
	b.Values[tmpslices[id][0]][tmpslices[id][1]] = 2
}
