package get

import (
	"context"
)

type UseCase interface {
	Execute(ctx context.Context) (string, error)
}

type Implementation struct {
}

func (useCase *Implementation) Execute(ctx context.Context) (string, error) {
	return "Pong", nil
}
