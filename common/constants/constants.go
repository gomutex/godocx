// Package constants provides constant values related to OpenXML namespaces, relationships, and other attributes.
package constants

import "encoding/xml"

var XMLHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")

const (
	OFFICE_DOC_TYPE    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	CORE_PROP_TYPE     = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
	EXTENDED_PROP_TYPE = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	StylesType         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
)

var (
	DrawingMLMainNS = "http://schemas.openxmlformats.org/drawingml/2006/main"
	DrawingMLPicNS  = "http://schemas.openxmlformats.org/drawingml/2006/picture"

	NameSpaceDocumentPropertiesVariantTypes = xml.Attr{Name: xml.Name{Local: "vt", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"}
	NameSpaceDrawing2016SVG                 = xml.Attr{Name: xml.Name{Local: "asvg", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2016/SVG/main"}
	NameSpaceDrawingML                      = xml.Attr{Name: xml.Name{Local: "a", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"}
	NameSpaceDrawingMLA14                   = xml.Attr{Name: xml.Name{Local: "a14", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2010/main"}
	NameSpaceDrawingMLChart                 = xml.Attr{Name: xml.Name{Local: "c", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"}
	NameSpaceDrawingMLSlicer                = xml.Attr{Name: xml.Name{Local: "sle", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2010/slicer"}
	NameSpaceDrawingMLSlicerX15             = xml.Attr{Name: xml.Name{Local: "sle15", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2012/slicer"}
	SourceRelationship                      = xml.Attr{Name: xml.Name{Local: "r", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"}
	SourceRelationshipChart20070802         = xml.Attr{Name: xml.Name{Local: "c14", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2007/8/2/chart"}
	SourceRelationshipChart2014             = xml.Attr{Name: xml.Name{Local: "c16", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2014/chart"}
	SourceRelationshipChart201506           = xml.Attr{Name: xml.Name{Local: "c16r2", Space: "xmlns"}, Value: "http://schemas.microsoft.com/office/drawing/2015/06/chart"}
	SourceRelationshipCompatibility         = xml.Attr{Name: xml.Name{Local: "mc", Space: "xmlns"}, Value: "http://schemas.openxmlformats.org/markup-compatibility/2006"}
)

const (
	XMLNS          = `http://schemas.openxmlformats.org/package/2006/relationships`
	HyperLinkStyle = "Hyperlink"
)

const (
	WMLNamespace    = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
	AltWMLNamespace = "http://purl.oclc.org/ooxml/wordprocessingml/main"

	WMLDrawingNS = "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
)

const (
	StrictNameSpaceDocumentPropertiesVariantTypes = "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes"
	StrictNameSpaceDrawingMLMain                  = "http://purl.oclc.org/ooxml/drawingml/main"
	StrictNameSpaceExtendedProperties             = "http://purl.oclc.org/ooxml/officeDocument/extendedProperties"
	StrictSourceRelationship                      = "http://purl.oclc.org/ooxml/officeDocument/relationships"
	StrictSourceRelationshipChart                 = "http://purl.oclc.org/ooxml/officeDocument/relationships/chart"
	StrictSourceRelationshipComments              = "http://purl.oclc.org/ooxml/officeDocument/relationships/comments"
	StrictSourceRelationshipExtendProperties      = "http://purl.oclc.org/ooxml/officeDocument/relationships/extendedProperties"
	StrictSourceRelationshipImage                 = "http://purl.oclc.org/ooxml/officeDocument/relationships/image"
	StrictSourceRelationshipOfficeDocument        = "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument"
)

const (
	NameSpaceDrawingMLMain                = "http://schemas.openxmlformats.org/drawingml/2006/main"
	NameSpaceDublinCore                   = "http://purl.org/dc/elements/1.1/"
	NameSpaceDublinCoreMetadataInitiative = "http://purl.org/dc/dcmitype/"
	NameSpaceDublinCoreTerms              = "http://purl.org/dc/terms/"
	NameSpaceExtendedProperties           = "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"
	NameSpaceXML                          = "http://www.w3.org/XML/1998/namespace"
	NameSpaceXMLSchemaInstance            = "http://www.w3.org/2001/XMLSchema-instance"
)

// Relationships
const (
	SourceRelationshipChart            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	SourceRelationshipComments         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
	SourceRelationshipExtendProperties = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	SourceRelationshipImage            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
	SourceRelationshipOfficeDocument   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	SourceRelationshipHyperLink        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
)

const (
	XMLNS_W = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XMLNS_R = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
)

const MediaPath = "word/media/"

const ConentTypeFileIdx = "[Content_Types].xml"
