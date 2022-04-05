package usecase

import (
	"context"
)

type DroneUsecase interface {
	RegisterDrone(ctx context.Context, request []byte) ([]byte, error)
}

type droneUsecase struct {
}

func NewDroneUsecase() DroneUsecase {
	return droneUsecase{}
}

func (d droneUsecase) RegisterDrone(ctx context.Context, request []byte) ([]byte, error) {
	return []byte{}, nil
}
