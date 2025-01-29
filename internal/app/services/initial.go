package services

import (
	"context"

	"go.uber.org/zap"
)

type InitialService struct {
	logger *zap.Logger
}

func NewInitialService(logger *zap.Logger) *InitialService {
	return &InitialService{
		logger: logger,
	}
}

func (i *InitialService) Initial(ctx context.Context) string {
	return "this is an example"
}
