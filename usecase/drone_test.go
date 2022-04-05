package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func Test_droneUsecase_RegisterDrone(t *testing.T) {
	type args struct {
		ctx     context.Context
		request []byte
	}
	tests := []struct {
		name    string
		d       droneUsecase
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "invalid serial number (contains numbers)",
			d:       droneUsecase{},
			args:    args{ctx: context.Background(), request: []byte(`{"serial_number": "SDFSDFDSf123213adsfasdfadsfasdf3","model":"Lightweight","weight":100.0}`)},
			want:    fmt.Sprintf("serial_number: %s does not validate as alpha", "SDFSDFDSf123213adsfasdfadsfasdf3"),
			wantErr: true,
		},
		{
			name:    "invalid serial number (more than 100)",
			d:       droneUsecase{},
			args:    args{ctx: context.Background(), request: []byte(`{"serial_number": "RepeatedTextRepeatedTextRepeatedTextRepeatedTextRepeatedTextRepeatedTextRepeatedTextRepeatedTextRepeatedTextRepeatedText","model":"Lightweight","weight":100.0}`)},
			want:    `serial_number: maximum length is 100`,
			wantErr: true,
		},
		{
			name:    "invalid weight (chars instead of float number)",
			d:       droneUsecase{},
			args:    args{ctx: context.Background(), request: []byte(`{"serial_number": "SDXDFSFEAADSF","model":"Lightweight","weight":"x"}`)},
			want:    `weight: invalid input , should be float`,
			wantErr: true,
		},
		{
			name:    "invalid model option",
			d:       droneUsecase{},
			args:    args{ctx: context.Background(), request: []byte(`{"serial_number": "SDXDFSFEAADSF","model":"x","weight":100.0}`)},
			want:    `model: should be option of [Lightweight Middleweight Cruiserweight Heavyweight]`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := droneUsecase{}
			_, err := d.RegisterDrone(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("droneUsecase.RegisterDrone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("droneUsecase.RegisterDrone() = %v, want %v", err.Error(), tt.want)
			}
		})
	}
}
