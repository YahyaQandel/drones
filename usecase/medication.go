package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"drones.com/iofile"
	"drones.com/repository"
)

const MEDICATION_IMAGES_PATH = "medications-images/"

type IMedicationeUsecase interface {
	RegisterMedication(ctx context.Context, r *http.Request) ([]byte, error)
}

type medicationUsecase struct {
	medicationRepo repository.IDroneRepo
	ioFile         iofile.IIOFile
}

func NewMedicationUsecase(droneRepository repository.IDroneRepo, ioFile iofile.IIOFile) IMedicationeUsecase {
	return medicationUsecase{medicationRepo: droneRepository, ioFile: ioFile}
}

func (d medicationUsecase) RegisterMedication(ctx context.Context, r *http.Request) (response []byte, err error) {
	r.ParseMultipartForm(32 * iofile.MB)
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize, fileType := d.ioFile.GetInfo(file, *handler)
	size := float64(fileSize) / float64(iofile.MB)
	if size >= 5 {
		return []byte{}, errors.New(fmt.Sprintf("image size '%f' is larger than 5mb", size))
	}
	if fileType != "image/jpeg" && fileType != "image/jpg" && fileType != "image/png" {
		return []byte{}, errors.New(fmt.Sprintf("unsupported file type %s, acceptable types (png/jpeg/jpg)", fileType))
	}
	err = d.ioFile.SaveImage(file, *handler, MEDICATION_IMAGES_PATH)
	if err != nil {
		return []byte{}, err
	}
	// repo should work here
	return nil, nil
}
