package opc

import (
	"fmt"
	"path"
)

func GetRelsURI(filePath string) (*string, error) {
	baseURI := path.Dir(filePath)
	_, fileName := path.Split(filePath)
	relsFilename := fmt.Sprintf("%s.rels", fileName)
	relsURI := path.Join(baseURI, "_rels", relsFilename)
	return &relsURI, nil
}
