package docx

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	types := ContentTypes{
		Default: []Default{
			{"xml", "application/xml"},
			{"rels", "application/vnd.openxmlformats-package.relationships+xml"},
			{"jpeg", "image/jpeg"},
		},
		Override: []Override{
			{"/word/document.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"},
			{"/customXml/itemProps1.xml", "application/vnd.openxmlformats-officedocument.customXmlProperties+xml"},
			{"/word/numbering.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"},
			{"/word/styles.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"},
			{"/word/stylesWithEffects.xml", "application/vnd.ms-word.stylesWithEffects+xml"},
			{"/word/settings.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"},
			{"/word/webSettings.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"},
			{"/word/fontTable.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"},
			{"/word/theme/theme1.xml", "application/vnd.openxmlformats-officedocument.theme+xml"},
			{"/docProps/core.xml", "application/vnd.openxmlformats-package.core-properties+xml"},
			{"/docProps/app.xml", "application/vnd.openxmlformats-officedocument.extended-properties+xml"},
		},
	}

	contentType, err := MIMEFromExt(".png")
	if err != nil {
		t.Error(err)
	}
	err = types.AddExtension("png", contentType)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	types.AddOverride("/customXml/item2.xml", "application/vnd.openxmlformats-officedocument.customXmlProperties+xml")

	xmlData, err := xml.Marshal(types)
	if err != nil {
		t.Fatalf("Error marshalling XML: %v", err)
	}

	expectedXML := `<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types"><Default Extension="xml" ContentType="application/xml"></Default><Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"></Default><Default Extension="jpeg" ContentType="image/jpeg"></Default><Default Extension="png" ContentType="image/png"></Default><Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"></Override><Override PartName="/customXml/itemProps1.xml" ContentType="application/vnd.openxmlformats-officedocument.customXmlProperties+xml"></Override><Override PartName="/word/numbering.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"></Override><Override PartName="/word/styles.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"></Override><Override PartName="/word/stylesWithEffects.xml" ContentType="application/vnd.ms-word.stylesWithEffects+xml"></Override><Override PartName="/word/settings.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"></Override><Override PartName="/word/webSettings.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"></Override><Override PartName="/word/fontTable.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"></Override><Override PartName="/word/theme/theme1.xml" ContentType="application/vnd.openxmlformats-officedocument.theme+xml"></Override><Override PartName="/docProps/core.xml" ContentType="application/vnd.openxmlformats-package.core-properties+xml"></Override><Override PartName="/docProps/app.xml" ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml"></Override><Override PartName="/customXml/item2.xml" ContentType="application/vnd.openxmlformats-officedocument.customXmlProperties+xml"></Override></Types>`

	if string(xmlData) != expectedXML {
		t.Errorf("Marshalled XML does not match expected. Got: %s, Expected: %s", string(xmlData), expectedXML)
	}
}

func TestUnmarshal(t *testing.T) {
	xmlData := `
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
	<Default Extension="xml" ContentType="application/xml"/>
	<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
	<Default Extension="jpeg" ContentType="image/jpeg"/>
	<Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
	<Override PartName="/customXml/itemProps1.xml" ContentType="application/vnd.openxmlformats-officedocument.customXmlProperties+xml"/>
	<Override PartName="/word/numbering.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"/>
	<Override PartName="/word/styles.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"/>
	<Override PartName="/word/stylesWithEffects.xml" ContentType="application/vnd.ms-word.stylesWithEffects+xml"/>
	<Override PartName="/word/settings.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"/>
	<Override PartName="/word/webSettings.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"/>
	<Override PartName="/word/fontTable.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"/>
	<Override PartName="/word/theme/theme1.xml" ContentType="application/vnd.openxmlformats-officedocument.theme+xml"/>
	<Override PartName="/docProps/core.xml" ContentType="application/vnd.openxmlformats-package.core-properties+xml"/>
	<Override PartName="/docProps/app.xml" ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml"/>
</Types>`

	var types ContentTypes
	err := xml.Unmarshal([]byte(xmlData), &types)
	if err != nil {
		t.Fatalf("Error unmarshalling XML: %v", err)
	}

	expected := ContentTypes{
		XMLName: xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/content-types", Local: "Types"},
		Default: []Default{
			{"xml", "application/xml"},
			{"rels", "application/vnd.openxmlformats-package.relationships+xml"},
			{"jpeg", "image/jpeg"},
		},
		Override: []Override{
			{"/word/document.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"},
			{"/customXml/itemProps1.xml", "application/vnd.openxmlformats-officedocument.customXmlProperties+xml"},
			{"/word/numbering.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"},
			{"/word/styles.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"},
			{"/word/stylesWithEffects.xml", "application/vnd.ms-word.stylesWithEffects+xml"},
			{"/word/settings.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"},
			{"/word/webSettings.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"},
			{"/word/fontTable.xml", "application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"},
			{"/word/theme/theme1.xml", "application/vnd.openxmlformats-officedocument.theme+xml"},
			{"/docProps/core.xml", "application/vnd.openxmlformats-package.core-properties+xml"},
			{"/docProps/app.xml", "application/vnd.openxmlformats-officedocument.extended-properties+xml"},
		},
	}

	if !reflect.DeepEqual(types, expected) {
		t.Errorf("Unmarshalled struct does not match expected. Got: %+v, Expected: %+v", types, expected)
	}
}

func TestAddExtension(t *testing.T) {
	types := ContentTypes{}
	mime, err := MIMEFromExt(".png")
	if err != nil {
		t.Error(err)
	}
	err = types.AddExtension("png", mime)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	expected := ContentTypes{
		Default: []Default{
			{"png", "image/png"},
		},
	}

	if !reflect.DeepEqual(types.Default, expected.Default) {
		t.Errorf("AddDefault did not add correctly. Got: %+v, Expected: %+v", types.Default, expected.Default)
	}
}

func TestAddOverride(t *testing.T) {
	types := ContentTypes{}
	types.AddOverride("/customXml/item2.xml", "application/vnd.openxmlformats-officedocument.customXmlProperties+xml")

	expected := ContentTypes{
		Override: []Override{
			{"/customXml/item2.xml", "application/vnd.openxmlformats-officedocument.customXmlProperties+xml"},
		},
	}

	if !reflect.DeepEqual(types.Override, expected.Override) {
		t.Errorf("AddOverride did not add correctly. Got: %+v, Expected: %+v", types.Override, expected.Override)
	}
}
