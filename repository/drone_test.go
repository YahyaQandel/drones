package repository

import (
	"context"
	"reflect"
	"testing"

	"drones.com/repository/db"
	"drones.com/repository/entity"
	"drones.com/testsetup"
	"gorm.io/gorm"
)

func TestDroneRepository_Create(t *testing.T) {
	dbClient, err := db.ConnectToDatabase(testsetup.TEST_CONNECTION_STRING)
	testdb := testsetup.DB{}
	ctx := context.Background()
	serialNumber := "XSfsdfdsfasdfsdfdksfsdf"
	drone := entity.Drone{
		SerialNumber:    serialNumber,
		Model:           "Lightweight",
		BatteryCapacity: 100,
		Weight:          50,
		State:           "IDLE",
	}
	if err != nil {
		t.Skipf("failed to connect to database: %v", err)
	}
	type fields struct {
		client *gorm.DB
	}
	type args struct {
		ctx   context.Context
		drone entity.Drone
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test create drone set to database",
			fields: fields{
				client: dbClient,
			},
			args: args{
				ctx:   ctx,
				drone: drone,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cdb := DroneRepository{
				client: tt.fields.client,
			}
			testdb.Init(dbClient)
			dbClient.Begin()
			defer dbClient.Rollback()
			_, err = cdb.Create(tt.args.ctx, tt.args.drone)
			if (err != nil) != tt.wantErr {
				t.Errorf("DroneRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			result := entity.Drone{}
			testdb.Get(context.Background(), tt.args.drone, &result)
			result.ID = 0 // to match the inserted object
			if !reflect.DeepEqual(tt.args.drone, result) {
				t.Errorf("DroneRepository.Create() error = %v, want %v", result, tt.args.drone)
			}
		})
	}
}
