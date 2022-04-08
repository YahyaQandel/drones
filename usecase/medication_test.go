package usecase

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	iomocks "drones.com/iofile/mocks"
	"drones.com/repository"
	"drones.com/repository/mocks"
)

func Test_medicationUsecase_RegisterMedication(t *testing.T) {
	type fields struct {
		medicationRepo repository.IDroneRepo
	}
	type args struct {
		ctx              context.Context
		r                *http.Request
		imageName        string
		medicationParams map[string]string
		medicationRepo   repository.IMedicationRepo
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse []byte
		wantErr      bool
		want         string
	}{
		{
			name:   "invalid image ext",
			fields: fields{},
			args: args{
				ctx:              context.Background(),
				imageName:        "invalid_image.txt",
				medicationParams: map[string]string{"name": "rx", "code": "RX_10", "weight": "10.1"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantResponse: []byte{},
			wantErr:      true,
			want:         "unsupported file type text/plain; charset=utf-8, acceptable types (png/jpeg/jpg)",
		},
		{
			name: "image larger than 5 mb",
			args: args{
				ctx:              context.Background(),
				imageName:        "image_5mb.jpg",
				medicationParams: map[string]string{"name": "rx", "code": "RX_10", "weight": "10.1"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantErr: true,
			want:    "image size '5.002336' is larger than 5mb",
		},
		{
			name: "successful image save",
			args: args{
				ctx:              context.Background(),
				imageName:        "test_image.jpeg",
				medicationParams: map[string]string{"name": "rx", "code": "RX_10", "weight": "10.1"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid medication name (special chars )",
			args: args{
				ctx:              context.Background(),
				imageName:        "test_image.jpeg",
				medicationParams: map[string]string{"name": "$$#xfsdf", "code": "RX_10", "weight": "10.1"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantErr: true,
			want:    "invalid medication name, (allowed only letters, numbers, '-', '_')",
		},
		{
			name: "invalid medication name ( empty )",
			args: args{
				ctx:              context.Background(),
				imageName:        "test_image.jpeg",
				medicationParams: map[string]string{"name": "", "code": "RX_10", "weight": "10.1"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantErr: true,
			want:    "invalid medication name, (allowed only letters, numbers, '-', '_')",
		},
		{
			name: "invalid medication code ( lowercase letters )",
			args: args{
				ctx:              context.Background(),
				imageName:        "test_image.jpeg",
				medicationParams: map[string]string{"name": "Rx", "code": "rx_10", "weight": "10.1"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantErr: true,
			want:    "invalid medication code, (allowed only upper case letters, underscore and numbers)",
		},
		{
			name: "invalid medication code ( empty )",
			args: args{
				ctx:              context.Background(),
				imageName:        "test_image.jpeg",
				medicationParams: map[string]string{"name": "Rx", "code": "", "weight": "10.1"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantErr: true,
			want:    "invalid medication code, (allowed only upper case letters, underscore and numbers)",
		},
		{
			name: "invalid weight",
			args: args{
				ctx:              context.Background(),
				imageName:        "test_image.jpeg",
				medicationParams: map[string]string{"name": "Rx", "code": "RX_10", "weight": "x"},
				medicationRepo:   mocks.NewMockedMedicationRepository(),
			},
			wantErr: true,
			want:    "invalid medication wegiht",
		},
		{
			name: "add medication twice",
			args: args{
				ctx:              context.Background(),
				imageName:        "test_image.jpeg",
				medicationParams: map[string]string{"name": "rx", "code": "RX", "weight": "10"},
				medicationRepo:   mocks.NewMedicationExistsRepository(),
			},
			wantErr: false,
			want:    "medication with same code already exists",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, w := createMultipartFormData(t, "image", fmt.Sprintf("tests-fixtures/%s", tt.args.imageName), tt.args.medicationParams)
			req, err := http.NewRequest("POST", "", &b)
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", w.FormDataContentType())
			tt.args.r = req
			medicationRepo := mocks.NewMockedMedicationRepository()
			mockedIOFile := iomocks.NewMockedIOFile()
			medicationUsecase := NewMedicationUsecase(medicationRepo, mockedIOFile)
			_, err = medicationUsecase.RegisterMedication(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("medicationUsecase.RegisterMedication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.want != err.Error() {
				t.Errorf("medicationUsecase.RegisterMedication() = %v, want %v", err.Error(), tt.want)
			}
		})
	}
}

func createMultipartFormData(t *testing.T, fieldName, fileName string, params map[string]string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file := mustOpen(fileName)
	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		t.Errorf("Error creating writer: %v", err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}
	for key, val := range params {
		_ = w.WriteField(key, val)
	}
	w.Close()
	return b, w
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println("PWD: ", pwd)
		panic(err)
	}
	return r
}
