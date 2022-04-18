package mocks

import (
	"context"

	"drones.com/repository"
)

type IMockedLogUsecase interface {
	LogDroneInfo(ctx context.Context) error
}

type MockedLogUsecase struct {
	droneRepo    repository.IDroneRepo
	droneLogRepo repository.IDroneLogRepo
}

func NewMockedLogUsecase(droneRepo repository.IDroneRepo, droneLogRepo repository.IDroneLogRepo) IMockedLogUsecase {
	return MockedLogUsecase{droneRepo: droneRepo, droneLogRepo: droneLogRepo}
}

func (l MockedLogUsecase) LogDroneInfo(ctx context.Context) error {
	return nil
}
