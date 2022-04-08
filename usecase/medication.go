package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"

	"drones.com/iofile"
	"drones.com/repository"
	repoEntity "drones.com/repository/entity"
)

type IMedicationeUsecase interface {
	RegisterMedication(ctx context.Context, r *http.Request) ([]byte, error)
}

type medicationUsecase struct {
	medicationRepo repository.IMedicationRepo
	ioFile         iofile.IIOFile
}

func NewMedicationUsecase(medicationRepository repository.IMedicationRepo, ioFile iofile.IIOFile) IMedicationeUsecase {
	return medicationUsecase{medicationRepo: medicationRepository, ioFile: ioFile}
}

func (d medicationUsecase) RegisterMedication(ctx context.Context, r *http.Request) (response []byte, err error) {
	medicationRequest := repoEntity.Medication{}
	err = d.fillInMedicationObjFromRequest(r, &medicationRequest)
	if err != nil {
		return []byte{}, err
	}
	medicationExists, err := d.medicationRepo.Get(ctx, medicationRequest)
	if !d.medicationRepo.IsNotFoundErr(err) && medicationExists.Code == medicationRequest.Code {
		return []byte{}, errors.New("medication with this code already exists")
	}
	medicationObject, err := d.medicationRepo.Create(ctx, medicationRequest)
	if err != nil {
		return []byte{}, err
	}
	imageName, err := d.ioFile.SaveImage(medicationRequest.Code)
	if err != nil {
		return []byte{}, err
	}
	medicationRequest.Image = imageName
	medicationJsonObj, err := json.Marshal(medicationObject)
	if err != nil {
		return []byte{}, err
	}
	return medicationJsonObj, nil
}

func getFormFieldValue(p *multipart.Part, field string) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(p)
	return buf.String()
}

func isValidMedicationName(name string) bool {
	// allow only letters , numbers '_' and -
	re := regexp.MustCompile("^[a-zA-Z0-9_-]+$")
	return re.MatchString(name)
}

func isValidMedicationCode(name string) bool {
	// allow only uppercase letters , numbers '_'
	re := regexp.MustCompile("^[A-Z0-9_]+$")
	return re.MatchString(name)
}

func isValidWeight(weight string) (isValid bool, result float64) {
	isValid = false
	result, err := strconv.ParseFloat(weight, 64)
	if err != nil {
		return
	}
	return true, result
}

func (d medicationUsecase) fillInMedicationObjFromRequest(r *http.Request, medication *repoEntity.Medication) error {
	reader, err := r.MultipartReader()
	if err != nil {
		return err
	}
	for {
		p, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		field := p.FormName()

		switch field {
		case "name":
			medicationName := getFormFieldValue(p, field)
			if !isValidMedicationName(medicationName) {
				return errors.New("invalid medication name, (allowed only letters, numbers, '-', '_')")
			}
			medication.Name = medicationName
		case "code":
			medicationCode := getFormFieldValue(p, field)
			if !isValidMedicationCode(medicationCode) {
				return errors.New("invalid medication code, (allowed only upper case letters, underscore and numbers)")
			}
			medication.Code = medicationCode
		case "image":
			fileSize, fileType, err := d.ioFile.GetInfo(p)
			if err != nil {
				return err
			}
			size := float64(fileSize) / float64(iofile.MB)
			if size >= 5 {
				return errors.New(fmt.Sprintf("image size '%f' is larger than 5mb", size))
			}
			if fileType != "image/jpeg" && fileType != "image/jpg" && fileType != "image/png" {
				return errors.New(fmt.Sprintf("unsupported file type %s, acceptable types (png/jpeg/jpg)", fileType))
			}
		case "weight":
			medicationWeight := getFormFieldValue(p, field)
			isValid, weightFloat := isValidWeight(medicationWeight)
			if !isValid {
				return errors.New("invalid medication wegiht")
			}
			medication.Weight = weightFloat
		}

	}
	return nil
}
