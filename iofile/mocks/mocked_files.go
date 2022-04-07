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

func (i mockedIOFile) SaveImage(file multipart.File, handler multipart.FileHeader, filePath string) (err error) {
	return nil
}

func (i mockedIOFile) GetInfo(file multipart.File, header multipart.FileHeader) (size int64, fileType string) {
	iofile := iofile.NewIOFile("")
	return iofile.GetInfo(file, header)
}
