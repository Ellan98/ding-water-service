package domain

import (
	"context"
	"fmt"
)

type Repository interface {
	Post(ctx context.Context, model string) (*User, error)
}

type NotFound struct {
	Model string
}

func (n NotFound) Error() string {
	return fmt.Sprintf("Wrong question %s asked", n.Model)
}
