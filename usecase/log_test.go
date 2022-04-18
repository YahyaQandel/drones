package usecase

import (
	"context"
	"reflect"
	"testing"

	"drones.com/repository"
	repoEntity "drones.com/repository/entity"
	"drones.com/repository/mocks"
	"drones.com/usecase/entity"
)

func Test_logUsecase_LogDroneInfo(t *testing.T) {
	type fields struct {
		droneRepo    repository.IDroneRepo
		droneLogRepo repository.IDroneLogRepo
	}
	type args struct {
		ctx          context.Context
		droneRepo    repository.IDroneRepo
		droneLogRepo repository.IDroneLogRepo
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		wantResponse repoEntity.DroneLog
	}{
		{
			name: "save single drone log",
			args: args{
				droneRepo:    mocks.NewMockedDroneExistsRepository(),
				droneLogRepo: mocks.NewMockedDroneLogExistsRepository(),
			},
			wantErr: false,
			wantResponse: repoEntity.DroneLog{
				DroneSerialNumber: "XDX",
				DroneBatteryLevel: 30,
				DroneState:        string(entity.LOADED),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logUsecase := NewLogUsecase(tt.args.droneRepo, tt.args.droneLogRepo)
			err := logUsecase.LogDroneInfo(tt.args.ctx)
			usecaseRepoDroneLog, err := tt.args.droneLogRepo.Get(tt.args.ctx, repoEntity.DroneLog{})
			if (err != nil) != tt.wantErr {
				t.Errorf("logUsecase.LogDroneInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(usecaseRepoDroneLog, tt.wantResponse) {
				t.Errorf("droneUsecase.RegisterDrone() = %v, want %v", usecaseRepoDroneLog, tt.wantResponse)
			}
		})
	}
}
