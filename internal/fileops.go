package internal

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

func ReadFileFromZip(file *zip.File) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}

	dat := make([]byte, 0, file.FileInfo().Size())
	buff := bytes.NewBuffer(dat)
	_, _ = io.Copy(buff, f)

	return buff.Bytes(), f.Close()
}

func FileToByte(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = f.Read(fileBytes)

	return fileBytes, nil
}
