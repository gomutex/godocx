package docx

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/gomutex/godocx/common/constants"
)

// CoreProperties represents the core properties of a document, such as title, creator, and version.
// It is used to store metadata information about the document.
type CoreProperties struct {
	Category       string
	ContentStatus  string
	Created        string
	Creator        string
	Description    string
	Identifier     string
	Keywords       string
	LastModifiedBy string
	Modified       string
	Revision       string
	Subject        string
	Title          string
	Language       string
	Version        string
}

// sourceCoreProps is an intermediate structure used for decoding XML data related to core properties.
type sourceCoreProps struct {
	XMLName        xml.Name       `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Title          string         `xml:"http://purl.org/dc/elements/1.1/ title,omitempty"`
	Subject        string         `xml:"http://purl.org/dc/elements/1.1/ subject,omitempty"`
	Creator        string         `xml:"http://purl.org/dc/elements/1.1/ creator"`
	Keywords       string         `xml:"keywords,omitempty"`
	Description    string         `xml:"http://purl.org/dc/elements/1.1/ description,omitempty"`
	LastModifiedBy string         `xml:"lastModifiedBy"`
	Language       string         `xml:"http://purl.org/dc/elements/1.1/ language,omitempty"`
	Identifier     string         `xml:"http://purl.org/dc/elements/1.1/ identifier,omitempty"`
	Revision       string         `xml:"revision,omitempty"`
	Created        *decodeDcTerms `xml:"http://purl.org/dc/terms/ created"`
	Modified       *decodeDcTerms `xml:"http://purl.org/dc/terms/ modified"`
	ContentStatus  string         `xml:"contentStatus,omitempty"`
	Category       string         `xml:"category,omitempty"`
	Version        string         `xml:"version,omitempty"`
}

// finalCoreProps is the final structure used for encoding core properties data to XML.
type finalCoreProps struct {
	FilePath       string       `xml:"-"`
	XMLName        xml.Name     `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Dc             string       `xml:"xmlns:dc,attr"`
	Dcterms        string       `xml:"xmlns:dcterms,attr"`
	Dcmitype       string       `xml:"xmlns:dcmitype,attr"`
	XSI            string       `xml:"xmlns:xsi,attr"`
	Title          string       `xml:"dc:title,omitempty"`
	Subject        string       `xml:"dc:subject,omitempty"`
	Creator        string       `xml:"dc:creator"`
	Keywords       string       `xml:"keywords,omitempty"`
	Description    string       `xml:"dc:description,omitempty"`
	LastModifiedBy string       `xml:"lastModifiedBy"`
	Language       string       `xml:"dc:language,omitempty"`
	Identifier     string       `xml:"dc:identifier,omitempty"`
	Revision       string       `xml:"revision,omitempty"`
	Created        *docxDcTerms `xml:"dcterms:created"`
	Modified       *docxDcTerms `xml:"dcterms:modified"`
	ContentStatus  string       `xml:"contentStatus,omitempty"`
	Category       string       `xml:"category,omitempty"`
	Version        string       `xml:"version,omitempty"`
}

// ExtendedProperties represents extended properties of a document, such as application details and statistics.
// It is used to store additional metadata information about the document.
type ExtendedProperties struct {
	Application       string
	ScaleCrop         bool
	DocSecurity       int
	Company           string
	LinksUpToDate     bool
	HyperlinksChanged bool
	AppVersion        string
}

// ctExtendedProperties is the structure used for encoding extended properties data to XML.
type ctExtendedProperties struct {
	FilePath             string         `xml:"-"`
	XMLName              xml.Name       `xml:"http://schemas.openxmlformats.org/officeDocument/2006/extended-properties Properties"`
	VT                   string         `xml:"xmlns:vt,attr"`
	Template             *string        `xml:"Template,omitempty"`
	TotalTime            *int           `xml:"TotalTime,omitempty"`
	Pages                *int           `xml:"Pages,omitempty"`
	Words                *int           `xml:"Words,omitempty"`
	Characters           *int           `xml:"Characters,omitempty"`
	Application          *string        `xml:"Application,omitempty"`
	DocSecurity          *int           `xml:"DocSecurity,omitempty"`
	Lines                *int           `xml:"Lines,omitempty"`
	Paragraphs           *int           `xml:"Paragraphs,omitempty"`
	ScaleCrop            *bool          `xml:"ScaleCrop,omitempty"`
	HeadingPairs         *HeadingPairs  `xml:"HeadingPairs,omitempty"`
	TitlesOfParts        *TitlesOfParts `xml:"TitlesOfParts,omitempty"`
	Manager              *string        `xml:"Manager,omitempty"`
	Company              *string        `xml:"Company,omitempty"`
	LinksUpToDate        *bool          `xml:"LinksUpToDate,omitempty"`
	CharactersWithSpaces *int           `xml:"CharactersWithSpaces,omitempty"`
	SharedDoc            *bool          `xml:"SharedDoc,omitempty"`
	HyperlinkBase        *string        `xml:"HyperlinkBase,omitempty"`
	HyperlinksChanged    *bool          `xml:"HyperlinksChanged,omitempty"`
	AppVersion           *string        `xml:"AppVersion,omitempty"`
}

// HeadingPairs represents a set of heading pairs used in extended properties.
type HeadingPairs struct {
	Vector Vector `xml:"HeadingPairs"`
}

// TitlesOfParts represents a set of titles of parts used in extended properties.
type TitlesOfParts struct {
	Vector Vector `xml:"TitlesOfParts"`
}

// Vector represents a vector structure used in extended properties.
type Vector struct {
	Size     int       `xml:"size,attr"`
	BaseType string    `xml:"baseType,attr"`
	Variant  []Variant `xml:"variant,omitempty"`
}

// Variant represents a variant structure used in extended properties.
type Variant struct {
	LPStr string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes lpstr,omitempty"`
	I4    int    `xml:"http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes i4,omitempty"`
}

// xmlNewDecoder creates a new XML decoder instance for the given reader.
// It returns the decoder, and it is used for decoding XML data.
func xmlNewDecoder(rdr io.Reader) (ret *xml.Decoder) {
	ret = xml.NewDecoder(rdr)
	// ret.CharsetReader = f.CharsetReader
	return
}

// LoadDocProps decodes the provided XML data and returns a CoreProperties instance.
// It is used to load core properties from the document file.
//
// Parameters:
//   - fileBytes: The XML data representing the core properties of the document.
//
// Returns:
//   - cp: The CoreProperties instance containing the decoded core properties.
//   - err: An error, if any occurred during the decoding process.
func LoadDocProps(fileBytes []byte) (cp *CoreProperties, err error) {
	core := new(sourceCoreProps)

	if err = xmlNewDecoder(bytes.NewReader(constants.TranslateNamespace(fileBytes))).
		Decode(core); err != nil && err != io.EOF {
		return
	}
	cp, err = &CoreProperties{
		Category:       core.Category,
		ContentStatus:  core.ContentStatus,
		Creator:        core.Creator,
		Description:    core.Description,
		Identifier:     core.Identifier,
		Keywords:       core.Keywords,
		LastModifiedBy: core.LastModifiedBy,
		Revision:       core.Revision,
		Subject:        core.Subject,
		Title:          core.Title,
		Language:       core.Language,
		Version:        core.Version,
	}, nil
	if core.Created != nil {
		cp.Created = core.Created.Text
	}
	if core.Modified != nil {
		cp.Modified = core.Modified.Text
	}
	return
}
