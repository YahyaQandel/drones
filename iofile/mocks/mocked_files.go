package mocks

import (
	"mime/multipart"

	"drones.com/iofile"
)

type mockedIOFile struct {
}

func NewMockedIOFile() iofile.IIOFile {
	return mockedIOFile{}
}

func (i mockedIOFile) SaveImage(filePath string) (imageName string, err error) {
	return "", nil
}

func (i mockedIOFile) GetInfo(part *multipart.Part) (int64, string, error) {
	iofile := iofile.NewIOFile("")
	return iofile.GetInfo(part)
}
