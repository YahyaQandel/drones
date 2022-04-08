package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"drones.com/repository"
	repoEntity "drones.com/repository/entity"
	"drones.com/repository/mocks"
	"drones.com/usecase/entity"
)

func Test_droneActionUsecase_LoadDrone(t *testing.T) {
	type fields struct {
		droneRepo repository.IDroneRepo
	}
	type args struct {
		ctx             context.Context
		request         []byte
		droneRepo       repository.IDroneRepo
		medicationRepo  repository.IMedicationRepo
		droneActionRepo repository.IDroneActionRepo
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse []byte
		wantErr      bool
		want         string
		droneState   string
	}{
		{
			name: "empty serial number and medication code validation",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "","medication_code":""}`),
				droneRepo:       mocks.NewMockedDroneRepository(),
				medicationRepo:  mocks.NewMockedMedicationRepository(),
				droneActionRepo: mocks.NewMockedDroneActionRepository(),
			},
			want:       fmt.Sprintf("drone_serial_number: non zero value required;medication_code: non zero value required"),
			wantErr:    true,
			droneState: string(entity.IDLE),
		},
		{
			name: "nonexists drone serial number validation",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX","medication_code":"RX"}`),
				droneRepo:       mocks.NewMockedDroneNotExistsRepository(),
				medicationRepo:  mocks.NewMockedMedicationRepository(),
				droneActionRepo: mocks.NewMockedDroneActionRepository(),
			},
			want:    fmt.Sprintf("drone not found with serial number 'XDX'"),
			wantErr: true,
		},
		{
			name: "nonexists medication code validation",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX","medication_code":"RX"}`),
				droneRepo:       mocks.NewMockedDroneRepository(),
				medicationRepo:  mocks.NewMockedMedicationNotExistsRepository(),
				droneActionRepo: mocks.NewMockedDroneActionRepository(),
			},
			want:       fmt.Sprintf("medication not found with code 'RX'"),
			wantErr:    true,
			droneState: string(entity.IDLE),
		},
		{
			// TODO: add test assertion that repo droneMedication create method not called
			name: "medication weight over than drone capacity validation",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX","medication_code":"RX"}`),
				droneRepo:       mocks.NewMockedDroneLessWeightThanMedicationRepository(50),
				medicationRepo:  mocks.NewMedicationOverWeightRepository(100),
				droneActionRepo: mocks.NewMockedDroneActionRepository(),
			},
			want:       fmt.Sprintf("drone weight '%0.2f' is not enough to carry medication with weight '%0.2f'", 50.00, 100.00),
			wantErr:    true,
			droneState: string(entity.IDLE),
		},
		{
			// TODO: add assertion that repo droneMedication create method not called
			name: "drone cannot be loaded if battery capacity is less than 25%",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX","medication_code":"RX"}`),
				droneRepo:       mocks.NewDroneBatteryLessThan25Repository(23),
				medicationRepo:  mocks.NewMockedMedicationRepository(),
				droneActionRepo: mocks.NewMockedDroneActionRepository(),
			},
			want:       fmt.Sprintf("cannot load drone with medication while battery capacity is '%0.2f'", 23.00),
			wantErr:    true,
			droneState: string(entity.IDLE),
		},
		{
			name: "drone loaded successfully",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX","medication_code":"RX"}`),
				droneRepo:       mocks.NewMockedLoadedDroneExistsRepository(),
				medicationRepo:  mocks.NewMedicationExistsRepository(),
				droneActionRepo: mocks.NewMockedDroneActionRepository(),
			},
			want:       fmt.Sprintf("drone '%s' loaded with medication '%s'", "XDX", "RX"),
			wantErr:    false,
			droneState: string(entity.LOADED),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneUsecase := NewDroneActionUsecase(tt.args.droneRepo, tt.args.medicationRepo, tt.args.droneActionRepo)
			_, err := droneUsecase.LoadDrone(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("droneActionUsecase.LoadDrone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("droneActionUsecase.LoadDrone() = %v, want %v", err.Error(), tt.want)
			}
			drone := entity.Drone{}
			err = json.Unmarshal(tt.args.request, &drone)
			droneRepo := repoEntity.Drone{}
			receivedDrone, err := tt.args.droneRepo.Get(tt.args.ctx, droneRepo)
			if tt.droneState != "" && receivedDrone.State != tt.droneState {
				t.Errorf("droneActionUsecase.LoadDrone() = %v, want %v", receivedDrone.State, tt.droneState)
			}
		})
	}
}
