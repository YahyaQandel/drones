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
		ctx       context.Context
		r         *http.Request
		imageName string
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
			name:    "invalid image ext",
			args:    args{ctx: context.Background(), imageName: "invalid_image.txt"},
			wantErr: true,
			want:    "unsupported file type application/octet-stream, acceptable types (png/jpeg/jpg)",
		},
		{
			name:    "image larger than 5 mb",
			args:    args{ctx: context.Background(), imageName: "image_5mb.jpg"},
			wantErr: true,
			want:    "image size '5.002336' is larger than 5mb",
		},
		{
			name:    "successful image save",
			args:    args{ctx: context.Background(), imageName: "test_image.jpeg"},
			wantErr: false,
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, w := createMultipartFormData(t, "image", fmt.Sprintf("tests-fixtures/%s", tt.args.imageName))
			req, err := http.NewRequest("POST", "", &b)
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", w.FormDataContentType())
			tt.args.r = req
			mockedRepo := mocks.NewMockedDroneRepository()
			mockedIOFile := iomocks.NewMockedIOFile()
			medicationUsecase := NewMedicationUsecase(mockedRepo, mockedIOFile)
			_, err = medicationUsecase.RegisterMedication(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("medicationUsecase.RegisterMedication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.want != err.Error() {
				t.Errorf("medicationUsecase.RegisterMedication() = %v, want %v", tt.want, err.Error())
			}
		})
	}
}

func createMultipartFormData(t *testing.T, fieldName, fileName string) (bytes.Buffer, *multipart.Writer) {
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
