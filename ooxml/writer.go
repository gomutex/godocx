package ooxml

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sort"
)

// Close method is used to close the RootDoc. Currently, it does not perform any specific actions.
func (rd *RootDoc) Close() error {
	var err error

	return err
}

// Write method writes the RootDoc to an io.Writer.
func (rd *RootDoc) Write(w io.Writer) error {
	_, err := rd.WriteTo(w)
	return err
}

// WriteTo implements io.WriterTo to write the RootDoc to an io.Writer.
func (rd *RootDoc) WriteTo(w io.Writer) (int64, error) {

	if err := rd.writeDirectToWriter(w); err != nil {
		return 0, err
	}
	return 0, nil
}

// writeDirectToWriter writes the RootDoc directly to an io.Writer using a zip.Writer.
func (rd *RootDoc) writeDirectToWriter(w io.Writer) error {
	zw := zip.NewWriter(w)
	if err := rd.writeToZip(zw); err != nil {
		_ = zw.Close()
		return err
	}
	return zw.Close()
}

// writeToZip provides a function to write to zip.Writer
func (rd *RootDoc) writeToZip(zw *zip.Writer) error {

	var (
		err   error
		files []string
	)

	docRelContent, err := marshal(rd.Document.DocRels)
	if err != nil {
		return err
	}
	rd.FileMap.Store(rd.Document.DocRels.RelativePath, []byte(docRelContent))

	rootRelContent, err := marshal(rd.RootRels)
	if err != nil {
		return err
	}
	rd.FileMap.Store(rd.RootRels.RelativePath, []byte(rootRelContent))

	docContent, err := marshal(rd.Document)
	if err != nil {
		return err
	}
	rd.FileMap.Store(rd.Document.relativePath, []byte(docContent))

	rd.FileMap.Range(func(path, content any) bool {
		files = append(files, path.(string))
		return true
	})

	sort.Strings(files)
	for _, path := range files {
		var fi io.Writer
		if fi, err = zw.Create(path); err != nil {
			break
		}
		content, _ := rd.FileMap.Load(path)
		_, err = fi.Write(content.([]byte))
	}

	return err
}

// Save method saves the RootDoc to the specified file path.
func (rd *RootDoc) Save() error {
	return rd.SaveTo(rd.Path)
}

// SaveTo method saves the RootDoc to the specified file path.
func (rd *RootDoc) SaveTo(fileName string) error {
	if fileName == "" {
		return errors.New("Destination file path is empty")
	}

	file, err := os.OpenFile(filepath.Clean(fileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return rd.Write(file)

}
