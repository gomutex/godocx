package docx

import (
	"encoding/xml"
	"errors"
	"strings"
)

// Types represents the root structure of the XML document.
type ContentTypes struct {
	XMLName  xml.Name   `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
	Default  []Default  `xml:"Default"`
	Override []Override `xml:"Override"`
}

// Default represents the <Default> XML element.
type Default struct {
	Extension   string `xml:"Extension,attr"`
	ContentType string `xml:"ContentType,attr"`
}

// Override represents the <Override> XML element.
type Override struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}

func (c *ContentTypes) AddExtension(extension, contentType string) error {
	c.Default = append(c.Default, Default{
		Extension:   extension,
		ContentType: contentType,
	})

	return nil
}

func (c *ContentTypes) AddOverride(partName, contentType string) error {
	c.Override = append(c.Override, Override{
		PartName:    partName,
		ContentType: contentType,
	})
	return nil
}

func MIMEFromExt(extension string) (string, error) {
	if strings.HasPrefix(extension, ".") {
		extension = strings.TrimPrefix(extension, ".")
	}

	switch extension {
	case "rels":
		return "application/vnd.openxmlformats-package.relationships+xml", nil
	case "xml":
		return "application/xml", nil
	case "jpg", "jpeg":
		return "image/jpeg", nil
	case "png":
		return "image/png", nil
	case "svg":
		return "image/svg+xml", nil
	case "gif":
		return "image/gif", nil
	case "bmp":
		return "image/bmp", nil
	case "tiff", "tif":
		return "image/tiff", nil
	case "docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document", nil
	case "xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", nil
	case "pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation", nil
	case "pdf":
		return "application/pdf", nil
	case "txt":
		return "text/plain", nil
	case "html", "htm":
		return "text/html", nil
	case "css":
		return "text/css", nil
	case "js":
		return "application/javascript", nil
	case "json":
		return "application/json", nil
	case "zip":
		return "application/zip", nil
	case "mp4":
		return "video/mp4", nil
	case "mp3":
		return "audio/mpeg", nil
	default:
		return "", errors.New("unsupported file extension")
	}
}
