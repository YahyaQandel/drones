package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"drones.com/repository"
	repoEntity "drones.com/repository/entity"
)

type ILogUsecase interface {
	LogDroneInfo(ctx context.Context) error
}

type logUsecase struct {
	droneRepo    repository.IDroneRepo
	droneLogRepo repository.IDroneLogRepo
}

func NewLogUsecase(droneRepo repository.IDroneRepo, droneLogRepo repository.IDroneLogRepo) ILogUsecase {
	return logUsecase{droneRepo: droneRepo, droneLogRepo: droneLogRepo}
}

func (l logUsecase) LogDroneInfo(ctx context.Context) error {
	allDrones, err := l.droneRepo.GetAll(ctx)
	if err != nil {
		return err
	}
	for _, drone := range allDrones {
		droneLog := repoEntity.DroneLog{
			DroneSerialNumber: drone.SerialNumber,
			DroneBatteryLevel: drone.BatteryCapacity,
			DroneState:        drone.State,
			CreatedTime:       time.Now(),
		}
		_, err := l.droneLogRepo.Create(ctx, droneLog)
		if err != nil {
			return err
		}
		log.Println(fmt.Sprintf("Drone `%s` log saved", drone.SerialNumber))
	}
	return nil
}
