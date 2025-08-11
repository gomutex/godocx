package docx

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/gomutex/godocx/common/constants"
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

	// Build a local deterministic snapshot rather than mutating rd.FileMap while writing
	snapshot := make(map[string][]byte)

	ct, err := marshal(rd.ContentType)
	if err != nil {
		return err
	}
	snapshot[constants.ConentTypeFileIdx] = []byte(ct)

	docRelContent, err := marshal(rd.Document.DocRels)
	if err != nil {
		return err
	}
	snapshot[rd.Document.DocRels.RelativePath] = docRelContent

	rootRelContent, err := marshal(rd.RootRels)
	if err != nil {
		return err
	}
	snapshot[rd.RootRels.RelativePath] = rootRelContent

	docContent, err := marshal(rd.Document)
	if err != nil {
		return err
	}
	snapshot[rd.Document.relativePath] = docContent

	docStyleBytes, err := marshal(rd.DocStyles)
	if err != nil {
		return err
	}
	snapshot[rd.DocStyles.RelativePath] = docStyleBytes

	// Persist numbering instances into numbering.xml if any
	if rd.Numbering != nil {
		// Apply numbering into a temporary buffer based on either existing or minimal content
		// by temporarily swapping root to a shallow copy with only the snapshot layer.
		// Minimal change: read existing numbering.xml if present from rd.FileMap, else let manager create it.
		if err := rd.Numbering.applyToFileMap(); err != nil {
			return err
		}
	}

	// Collect files from original map
	rd.FileMap.Range(func(path, content any) bool {
		p := path.(string)
		if _, exists := snapshot[p]; !exists {
			snapshot[p] = content.([]byte)
		}
		return true
	})

	// Now gather list of paths from snapshot
	for p := range snapshot {
		files = append(files, p)
	}

	sort.Strings(files)
	for _, path := range files {
		// Use a deterministic timestamp for reproducible archives
		hdr := &zip.FileHeader{
			Name:     path,
			Method:   zip.Deflate,
			Modified: time.Unix(0, 0).UTC(),
		}
		var fi io.Writer
		if fi, err = zw.CreateHeader(hdr); err != nil {
			break
		}
		_, err = fi.Write(snapshot[path])
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
