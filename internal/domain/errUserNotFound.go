package domain

import "fmt"

type ErrUserNotFound struct {
	ID uint64 `json:"id"`
}

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("user with ID %d not found", e.ID)
}
