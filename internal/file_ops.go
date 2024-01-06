package internal

import (
	"archive/zip"
	"bytes"
	"io"
)

func ReadFile(file *zip.File) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}

	dat := make([]byte, 0, file.FileInfo().Size())
	buff := bytes.NewBuffer(dat)
	_, _ = io.Copy(buff, f)

	return buff.Bytes(), f.Close()
}
