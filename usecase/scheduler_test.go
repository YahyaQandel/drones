package usecase

import (
	"context"
	"testing"

	"drones.com/repository"
	repoEntity "drones.com/repository/entity"
	"drones.com/repository/mocks"
)

func Test_schedulerUsecase_UpdateLoadedDronesBatteryLevel(t *testing.T) {
	medicationRepo := mocks.NewMockedMedicationRepository()
	droneActionRepo := mocks.NewMockedDroneActionRepository()
	type fields struct {
		droneActionUsecase IDroneActionUsecase
	}
	type args struct {
		droneRepo       repository.IDroneRepo
		mdicationRepo   repository.IMedicationRepo
		droneActionRepo repository.IDroneActionRepo
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantErr          bool
		wantBatteryLevel float64
	}{
		{
			name: "successful update drone battery level after 3 seconds if loaded",
			args: args{
				droneRepo: mocks.NewMockedUpdatedLoadedDronePeriodicallyRepository(repoEntity.Drone{
					BatteryCapacity: 50,
				}),
			},
			wantErr:          false,
			wantBatteryLevel: 47,
		},
		{
			name: "will not update drone battery level if it is already below 3",
			args: args{
				droneRepo: mocks.NewMockedUpdatedLoadedDronePeriodicallyRepository(repoEntity.Drone{
					BatteryCapacity: 2,
				}),
			},
			wantErr:          false,
			wantBatteryLevel: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneActionUsecase := NewDroneActionUsecase(tt.args.droneRepo, medicationRepo, droneActionRepo)
			schedulerUsecase := NewSchedulerUsecase(droneActionUsecase, tt.args.droneRepo)
			err := schedulerUsecase.UpdateLoadedDronesBatteryLevel(context.Background())
			droneAfterSchedulerrRun, err := tt.args.droneRepo.Get(context.Background(), repoEntity.Drone{})
			if err != nil {
				t.Errorf("schedulerUsecase.UpdateLoadedDronesBatteryLevel() error = %v", err)
			}
			if droneAfterSchedulerrRun.BatteryCapacity != tt.wantBatteryLevel {
				t.Errorf("schedulerUsecase.UpdateLoadedDronesBatteryLevel() error expected battery level = %v, wantErr %v", droneAfterSchedulerrRun.BatteryCapacity, tt.wantBatteryLevel)
			}
		})
	}
}
