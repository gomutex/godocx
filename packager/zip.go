package packager

import (
	"archive/zip"
	"fmt"
	"io"
)

// readZipFile reads the contents of a zip file and returns it as a byte slice.
// It returns an error if the file is not a valid zip file or if there is an error
// reading the file.
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
