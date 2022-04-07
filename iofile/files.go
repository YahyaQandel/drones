package iofile

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	MB = 1 << 20
)

type Sizer interface {
	Size() int64
}
type IIOFile interface {
	SaveImage(multipart.File, multipart.FileHeader, string) error
	GetInfo(multipart.File, multipart.FileHeader) (int64, string)
}

type IOFile struct {
	folderPath string
}

func NewIOFile(folderPath string) IIOFile {
	return IOFile{folderPath: folderPath}
}

func (i IOFile) SaveImage(file multipart.File, handler multipart.FileHeader, filePath string) (err error) {
	defer file.Close()
	f, err := os.OpenFile(i.folderPath+filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}

func (i IOFile) GetInfo(file multipart.File, header multipart.FileHeader) (size int64, fileType string) {
	fileHeader := make([]byte, 512)
	if _, err := file.Read(fileHeader); err != nil {
		return 0, ""
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		return 0, ""
	}
	size = file.(Sizer).Size()
	fileType = http.DetectContentType(fileHeader)
	return
}
