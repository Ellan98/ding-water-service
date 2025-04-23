package domain

import (
	"context"
	"fmt"
)

type Repository interface {
	Get(ctx context.Context, problem string) (*User, error)
}

type NotFound struct {
	Problem string
}

func (n NotFound) Error() string {
	return fmt.Sprintf("Wrong question %s asked", n.Problem)
}
