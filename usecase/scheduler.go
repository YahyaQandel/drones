package usecase

import (
	"context"
	"fmt"
	"log"

	"drones.com/repository"
)

type ISchedulerUsecase interface {
	UpdateLoadedDronesBatteryLevel(ctx context.Context) error
}

type schedulerUsecase struct {
	droneActionUsecase IDroneActionUsecase
	droneRepo          repository.IDroneRepo
}

func NewSchedulerUsecase(droneActionUsecase IDroneActionUsecase, droneRepo repository.IDroneRepo) ISchedulerUsecase {
	return schedulerUsecase{droneActionUsecase: droneActionUsecase, droneRepo: droneRepo}
}

func (s schedulerUsecase) UpdateLoadedDronesBatteryLevel(ctx context.Context) error {
	drones, err := s.droneActionUsecase.GetLoadedDrones(ctx)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	for _, drone := range drones {
		if drone.BatteryCapacity <= 3 {
			drone.BatteryCapacity = 0

		} else {
			drone.BatteryCapacity -= 3
		}
		_, err := s.droneRepo.Update(ctx, drone)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		log.Println(fmt.Sprintf("Drone %s battery level is %f", drone.SerialNumber, drone.BatteryCapacity))
	}
	return nil
}
