package iofile

import (
	"errors"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"os"
	"strings"
)

const (
	MB = 1 << 20
)

type Sizer interface {
	Size() int64
}
type IIOFile interface {
	SaveImage(filePath string) (imageName string, err error)
	GetInfo(part *multipart.Part) (int64, string, error)
}

type IOFile struct {
	folderPath string
	fileBytes  []byte
	fileExt    string
}

func NewIOFile(folderPath string) IIOFile {
	return &IOFile{folderPath: folderPath}
}

func (i *IOFile) SaveImage(filePath string) (imageName string, err error) {
	imageName = filePath + i.fileExt
	f, err := os.Create(i.folderPath + filePath + i.fileExt)
	_, err = f.Write(i.fileBytes)
	err = f.Close()
	if err != nil {
		return "", err
	}
	return
}

func (i *IOFile) GetInfo(part *multipart.Part) (size int64, fileType string, err error) {
	fileBytes, err := ioutil.ReadAll(part)
	if err != nil {
		return -1, "", errors.New("failed to read content of the part")
	}
	fname := part.FileName()
	dotIndex := strings.LastIndex(fname, ".")
	i.fileExt = fname[dotIndex:]
	fileType = mime.TypeByExtension(i.fileExt)
	i.fileBytes = fileBytes
	return int64(len(fileBytes)), fileType, nil
}
