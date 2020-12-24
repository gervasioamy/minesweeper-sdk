package sdk

import "fmt"

// MinesweeperError custom error for 4xx responses
type MinesweeperError struct {
	errorCode int
	message   string
}

func (e *MinesweeperError) Error() string {
	return fmt.Sprintf("%d - %s", e.errorCode, e.message)
}
