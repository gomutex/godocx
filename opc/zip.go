package opc

import (
	"archive/zip"
	"fmt"
	"io"
)

func readZipFile(zf *zip.File) ([]byte, error) {
	if zf == nil {
		return nil, fmt.Errorf("Invalid Zip file")
	}
	reader, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}
