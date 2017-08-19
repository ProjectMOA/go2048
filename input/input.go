package input

import (
	"fmt"

	"github.com/ProjectMOA/go2048/board"
	"github.com/eiannone/keyboard"
)

// Get returns a direction from the keyboard arrow keys. This call blocks the
// execution until either an arrow key or the ESC key are pressed.
func Get() (board.Direction, error) {
	for true {
		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			return -1, fmt.Errorf("error getting input key: %v", err)
		}
		switch key {
		case '￫':
			// Go left
			return board.Left, nil
		case '￬':
			// Go down
			return board.Down, nil
		case '￭':
			// Go up
			return board.Up, nil
		case '￪':
			// Go right
			return board.Right, nil
		case keyboard.Key(27):
			return -1, fmt.Errorf("Exit")
		}
	}
	return -1, fmt.Errorf("Unexpected loop break")
}
