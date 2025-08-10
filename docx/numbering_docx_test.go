package docx_test

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	godocx "github.com/gomutex/godocx"
)

// TestGenerateTestDocxWithLists creates testdata/numbering.docx containing multiple
// ordered, unordered, and nested lists to visually verify numbering behavior and
// programmatically confirm numbering instances are written.
func TestGenerateNumberingDocxSample(t *testing.T) {
	doc, err := godocx.NewDocument()
	if err != nil {
		t.Fatalf("NewDocument error: %v", err)
	}

	rd := doc

	// Ordered list instance A (abstract 1) with deeper nesting
	ordA := rd.NewListInstance(1)
	p := rd.AddParagraph("Ordered A 1")
	p.Numbering(ordA, 0)
	p = rd.AddParagraph("Nested A 1.1")
	p.Numbering(ordA, 1)
	p = rd.AddParagraph("Nested A 1.2")
	p.Numbering(ordA, 1)
	p = rd.AddParagraph("Nested A 1.2.i")
	p.Numbering(ordA, 2)
	p = rd.AddParagraph("Nested A 1.2.ii")
	p.Numbering(ordA, 2)
	p = rd.AddParagraph("Nested A 1.2.ii.A")
	p.Numbering(ordA, 3)
	p = rd.AddParagraph("Ordered A 2")
	p.Numbering(ordA, 0)

	// Ordered list instance B (same abstract 1, numbering should reset)
	ordB := rd.NewListInstance(1)
	p = rd.AddParagraph("Ordered B 1")
	p.Numbering(ordB, 0)
	p = rd.AddParagraph("Ordered B 2")
	p.Numbering(ordB, 0)

	// Unordered list instance C (abstract 2 bullets) with deeper nesting
	bulC := rd.NewListInstance(2)
	p = rd.AddParagraph("Bullet C 1")
	p.Numbering(bulC, 0)
	p = rd.AddParagraph("Nested C 1.1")
	p.Numbering(bulC, 1)
	p = rd.AddParagraph("Nested C 1.2")
	p.Numbering(bulC, 1)
	p = rd.AddParagraph("Nested C 1.2.i")
	p.Numbering(bulC, 2)
	p = rd.AddParagraph("Nested C 1.2.ii")
	p.Numbering(bulC, 2)
	p = rd.AddParagraph("Nested C 1.2.ii.A")
	p.Numbering(bulC, 3)
	p = rd.AddParagraph("Bullet C 2")
	p.Numbering(bulC, 0)

	// Unordered list instance D (same abstract 2, numbering should be independent)
	bulD := rd.NewListInstance(2)
	p = rd.AddParagraph("Bullet D 1")
	p.Numbering(bulD, 0)
	p = rd.AddParagraph("Bullet D 2")
	p.Numbering(bulD, 0)

	if err := os.MkdirAll(filepath.Join("..", "testdata"), 0o755); err != nil {
		t.Fatalf("mkdir testdata: %v", err)
	}
	outPath := filepath.Join("..", "testdata", "numbering.docx")
	if err := rd.SaveTo(outPath); err != nil {
		t.Fatalf("SaveTo error: %v", err)
	}

	// Read numbering.xml back from the saved .docx and verify instances exist
	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("reading saved docx: %v", err)
	}

	zr, err := zip.NewReader(bytes.NewReader(content), int64(len(content)))
	if err != nil {
		t.Fatalf("open zip: %v", err)
	}

	var numberingXML, documentXML string
	for _, f := range zr.File {
		if f.Name == "word/numbering.xml" || f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				t.Fatalf("open %s: %v", f.Name, err)
			}
			b, _ := io.ReadAll(rc)
			_ = rc.Close()
			if f.Name == "word/numbering.xml" {
				numberingXML = string(b)
			} else {
				documentXML = string(b)
			}
		}
	}

	if numberingXML == "" {
		t.Fatalf("word/numbering.xml not found in saved document")
	}

	// Verify that separate num instances exist and map to our multilevel abstracts
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(ordA)+`"`) ||
		!strings.Contains(numberingXML, `w:abstractNumId w:val="201"`) {
		t.Fatalf("numbering.xml missing ordered instance A (201)")
	}
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(ordB)+`"`) ||
		!strings.Contains(numberingXML, `w:abstractNumId w:val="201"`) {
		t.Fatalf("numbering.xml missing ordered instance B (201)")
	}
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(bulC)+`"`) ||
		!strings.Contains(numberingXML, `w:abstractNumId w:val="202"`) {
		t.Fatalf("numbering.xml missing bullet instance C (202)")
	}
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(bulD)+`"`) ||
		!strings.Contains(numberingXML, `w:abstractNumId w:val="202"`) {
		t.Fatalf("numbering.xml missing bullet instance D (202)")
	}

	// Verify restart override for ordered instances
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(ordA)+`"`) ||
		!strings.Contains(numberingXML, `<w:startOverride w:val="1"/>`) {
		t.Fatalf("ordered instance A missing startOverride")
	}
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(ordB)+`"`) ||
		!strings.Contains(numberingXML, `<w:startOverride w:val="1"/>`) {
		t.Fatalf("ordered instance B missing startOverride")
	}

	// Verify multilevel formats for ordered abstract 201
	if !strings.Contains(numberingXML, `<w:abstractNum w:abstractNumId="201"`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="0"><w:start w:val="1"/><w:numFmt w:val="decimal"/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="1"><w:start w:val="1"/><w:numFmt w:val="lowerLetter"/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="2"><w:start w:val="1"/><w:numFmt w:val="lowerRoman"/>`) {
		t.Fatalf("ordered abstract 201 missing expected numFmt per level")
	}

	// Verify bullet glyphs/fonts for abstract 202 (levels 0..3)
	if !strings.Contains(numberingXML, `<w:abstractNum w:abstractNumId="202"`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="0"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val=""/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="1"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val="○"/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="2"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val="■"/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="3"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val="♦"/>`) {
		t.Fatalf("bullet abstract 202 missing expected glyphs per level")
	}

	// Basic document.xml checks for nested levels
	if documentXML == "" {
		t.Fatalf("word/document.xml not found in saved document")
	}
	if !strings.Contains(documentXML, `<w:numId w:val="`+strconv.Itoa(ordA)+`"`) {
		t.Fatalf("document.xml missing reference to ordered instance A")
	}
	if !strings.Contains(documentXML, `<w:ilvl w:val="1"`) || !strings.Contains(documentXML, `<w:ilvl w:val="2"`) || !strings.Contains(documentXML, `<w:ilvl w:val="3"`) {
		t.Fatalf("document.xml missing nested level paragraphs (ilvl=1/2/3)")
	}
}
