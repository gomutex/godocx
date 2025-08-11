package docx

import (
	"strconv"
	"strings"
	"testing"
)

func TestNewListInstance(t *testing.T) {
	root := NewRootDoc()

	// Test creating list instances
	numId1 := root.NewListInstance(1)
	numId2 := root.NewListInstance(2)
	numId3 := root.NewListInstance(1) // Same abstract ID, different instance

	// Verify that unique numIds are returned
	if numId1 == numId2 {
		t.Errorf("Expected different numIds, got %d and %d", numId1, numId2)
	}
	if numId1 == numId3 {
		t.Errorf("Expected different numIds for same abstract ID, got %d and %d", numId1, numId3)
	}
	if numId2 == numId3 {
		t.Errorf("Expected different numIds, got %d and %d", numId2, numId3)
	}

	// Test that numIds are sequential
	if numId1 != 1 || numId2 != 2 || numId3 != 3 {
		t.Errorf("Expected sequential numIds starting from 1, got %d, %d, %d", numId1, numId2, numId3)
	}
}

func TestGetNumberingXML(t *testing.T) {
	root := NewRootDoc()

	// Create list instances
	_ = root.NewListInstance(1)
	_ = root.NewListInstance(2)

	// Get XML
	xml, err := root.Numbering.GetNumberingXML()
	if err != nil {
		t.Fatalf("Error getting numbering XML: %v", err)
	}

	xmlStr := string(xml)

	// Verify the XML contains expected elements
	if !strings.Contains(xmlStr, "w:numbering") {
		t.Error("XML should contain w:numbering element")
	}
	if !strings.Contains(xmlStr, `w:numId="1"`) {
		t.Error("XML should contain first numId")
	}
	if !strings.Contains(xmlStr, `w:numId="2"`) {
		t.Error("XML should contain second numId")
	}
	if !strings.Contains(xmlStr, `w:val="201"`) {
		t.Error("XML should contain first abstractNumId (decimal multilevel)")
	}
	if !strings.Contains(xmlStr, `w:val="202"`) {
		t.Error("XML should contain second abstractNumId (bullet multilevel)")
	}

	t.Logf("Generated XML: %s", xmlStr)
}

func TestNumberingResetsBetweenInstancesSameAbstract(t *testing.T) {
	root := NewRootDoc()

	// Two instances using the same abstract numbering definition should receive different numIds
	id1 := root.NewListInstance(1)
	id2 := root.NewListInstance(1)

	if id1 == id2 {
		t.Fatalf("Expected different numIds for separate instances, got %d and %d", id1, id2)
	}

	// Attach numbering to paragraphs to ensure they reference the correct instance ids
	p1 := &Paragraph{root: root}
	p1.Numbering(id1, 0)
	if p1.ct.Property == nil || p1.ct.Property.NumProp == nil || p1.ct.Property.NumProp.NumID == nil {
		t.Fatalf("Paragraph 1 numbering not set")
	}
	if got := p1.ct.Property.NumProp.NumID.Val; got != id1 {
		t.Fatalf("Paragraph 1 expected NumID %d, got %d", id1, got)
	}

	p2 := &Paragraph{root: root}
	p2.Numbering(id2, 0)
	if p2.ct.Property == nil || p2.ct.Property.NumProp == nil || p2.ct.Property.NumProp.NumID == nil {
		t.Fatalf("Paragraph 2 numbering not set")
	}
	if got := p2.ct.Property.NumProp.NumID.Val; got != id2 {
		t.Fatalf("Paragraph 2 expected NumID %d, got %d", id2, got)
	}
}

func TestEmptyNumberingXML(t *testing.T) {
	root := NewRootDoc()

	// Get XML - should be empty
	xml, err := root.Numbering.GetNumberingXML()
	if err != nil {
		t.Fatalf("Error getting numbering XML: %v", err)
	}

	if len(xml) != 0 {
		t.Errorf("Expected empty XML for no numbering instances, got: %s", string(xml))
	}
}

func TestNumbering_MultilevelAndBulletsXML(t *testing.T) {
	rd := NewRootDoc()

	// Ordered A (abstract 1 → 201)
	ordA := rd.NewListInstance(1)
	(&Paragraph{root: rd}).Numbering(ordA, 0)
	(&Paragraph{root: rd}).Numbering(ordA, 1)
	(&Paragraph{root: rd}).Numbering(ordA, 1)
	(&Paragraph{root: rd}).Numbering(ordA, 2)
	(&Paragraph{root: rd}).Numbering(ordA, 2)
	(&Paragraph{root: rd}).Numbering(ordA, 3)
	(&Paragraph{root: rd}).Numbering(ordA, 0)

	// Ordered B (reset numbering)
	ordB := rd.NewListInstance(1)
	(&Paragraph{root: rd}).Numbering(ordB, 0)
	(&Paragraph{root: rd}).Numbering(ordB, 0)

	// Bullets C (abstract 2 → 202)
	bulC := rd.NewListInstance(2)
	(&Paragraph{root: rd}).Numbering(bulC, 0)
	(&Paragraph{root: rd}).Numbering(bulC, 1)
	(&Paragraph{root: rd}).Numbering(bulC, 1)
	(&Paragraph{root: rd}).Numbering(bulC, 2)
	(&Paragraph{root: rd}).Numbering(bulC, 2)
	(&Paragraph{root: rd}).Numbering(bulC, 3)
	(&Paragraph{root: rd}).Numbering(bulC, 0)

	// Generate numbering.xml in the FileMap
	if err := rd.Numbering.applyToFileMap(); err != nil {
		t.Fatalf("apply numbering: %v", err)
	}
	v, ok := rd.FileMap.Load("word/numbering.xml")
	if !ok {
		t.Fatalf("word/numbering.xml not generated")
	}
	numberingXML := string(v.([]byte))

	// Instances map to multilevel abstracts
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(ordA)+`"`) || !strings.Contains(numberingXML, `w:abstractNumId w:val="201"`) {
		t.Fatalf("ordered instance A missing or not mapped to 201: %s", numberingXML)
	}
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(ordB)+`"`) || !strings.Contains(numberingXML, `w:abstractNumId w:val="201"`) {
		t.Fatalf("ordered instance B missing or not mapped to 201")
	}
	if !strings.Contains(numberingXML, `w:num w:numId="`+strconv.Itoa(bulC)+`"`) || !strings.Contains(numberingXML, `w:abstractNumId w:val="202"`) {
		t.Fatalf("bullet instance C missing or not mapped to 202")
	}

	// Restart overrides exist
	if !strings.Contains(numberingXML, `<w:startOverride w:val="1"/>`) {
		t.Fatalf("startOverride missing for instances")
	}

	// Ordered abstract formats
	if !strings.Contains(numberingXML, `<w:abstractNum w:abstractNumId="201"`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="1"><w:start w:val="1"/><w:numFmt w:val="lowerLetter"/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="2"><w:start w:val="1"/><w:numFmt w:val="lowerRoman"/>`) {
		t.Fatalf("ordered abstract 201 missing expected formats")
	}

	// Bullet glyphs (disc, hollow circle, square, diamond)
	if !strings.Contains(numberingXML, `<w:abstractNum w:abstractNumId="202"`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="0"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val=""/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="1"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val="○"/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="2"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val="■"/>`) ||
		!strings.Contains(numberingXML, `<w:lvl w:ilvl="3"><w:start w:val="1"/><w:numFmt w:val="bullet"/><w:lvlText w:val="♦"/>`) {
		t.Fatalf("bullet abstract 202 missing expected glyphs")
	}
}
