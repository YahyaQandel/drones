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
	usecaseEntity "drones.com/usecase/entity"
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
			// TODO: add assertion that repo droneMedication create method not called
			name: "drone cannot be loaded with same medication twice",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX","medication_code":"RX"}`),
				droneRepo:       mocks.NewMockedLoadedDroneRepository(),
				medicationRepo:  mocks.NewMockedMedicationRepository(),
				droneActionRepo: mocks.NewMockedDroneActionAlreadyLoadedWithSameDroneAndMedicationRepository(),
			},
			want:       fmt.Sprintf("drone with serial number 'XDX' is already loaded with medication with code 'RX'"),
			wantErr:    true,
			droneState: string(entity.LOADED),
		},
		{
			name: "drone loaded successfully",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX","medication_code":"RX"}`),
				droneRepo:       mocks.NewMockedDroneExistsRepository(),
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

func Test_droneActionUsecase_GetLoadedMedicationItems(t *testing.T) {
	type fields struct {
		droneRepo           repository.IDroneRepo
		medicationRepo      repository.IMedicationRepo
		droneMedicationRepo repository.IDroneActionRepo
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
		wantResponse []repoEntity.Medication
		wantErr      bool
	}{
		{
			name: "get all loaded medications for specific drone",
			args: args{
				ctx:             context.Background(),
				request:         []byte(`{"drone_serial_number": "XDX"}`),
				droneRepo:       mocks.NewMockedDroneRepository(),
				medicationRepo:  mocks.NewMedicationGetAllRepository(),
				droneActionRepo: mocks.NewMockedDroneActionGetAllMedicationRepository(),
			},
			wantResponse: []repoEntity.Medication{
				{
					ID:     1,
					Name:   "Aspirin",
					Code:   "RX",
					Weight: 100,
					Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
				},
				{
					ID:     2,
					Name:   "Advil",
					Code:   "UX",
					Weight: 200,
					Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
				},
				{
					ID:     3,
					Name:   "Vicodin",
					Code:   "DX",
					Weight: 300,
					Image:  "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneActionUsecase := NewDroneActionUsecase(tt.args.droneRepo, tt.args.medicationRepo, tt.args.droneActionRepo)
			gotResponse, err := droneActionUsecase.GetLoadedMedicationItems(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("droneActionUsecase.GetLoadedMedicationItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("droneActionUsecase.GetLoadedMedicationItems() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func Test_droneActionUsecase_GetAvailableDrones(t *testing.T) {
	type fields struct {
		droneRepo       repository.IDroneRepo
		medicationRepo  repository.IMedicationRepo
		droneActionRepo repository.IDroneActionRepo
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse []repoEntity.Drone
		wantErr      bool
	}{
		{
			name: "get all available drones",
			args: args{context.Background()},
			fields: fields{
				droneRepo:       mocks.NewMockedDroneGetAllAvailableRepository(),
				medicationRepo:  mocks.NewMockedMedicationRepository(),
				droneActionRepo: mocks.NewMockedDroneActionRepository(),
			},
			wantResponse: []repoEntity.Drone{
				{
					ID:              1,
					SerialNumber:    "12345",
					State:           string(usecaseEntity.IDLE),
					Model:           "LightWeight",
					BatteryCapacity: 60,
					Weight:          10,
				},
				{
					ID:              2,
					SerialNumber:    "54321",
					State:           string(usecaseEntity.IDLE),
					Model:           "HeavyWeight",
					BatteryCapacity: 80,
					Weight:          20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			droneActionUsecase := NewDroneActionUsecase(tt.fields.droneRepo, tt.fields.medicationRepo, tt.fields.droneActionRepo)
			gotResponse, err := droneActionUsecase.GetAvailableDrones(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("droneActionUsecase.GetAvailableDrones() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("droneActionUsecase.GetAvailableDrones() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
