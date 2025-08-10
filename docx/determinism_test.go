package docx_test

import (
	"bytes"
	"testing"

	godocx "github.com/gomutex/godocx"
	docxpkg "github.com/gomutex/godocx/docx"
	"github.com/stretchr/testify/require"
)

// Ensures that writing the same document multiple times yields identical bytes.
func TestDeterministicWrite_SameInstance(t *testing.T) {
	rd := buildElaborateDoc(t)

	var b1, b2 bytes.Buffer
	require.NoError(t, rd.Write(&b1))
	require.NoError(t, rd.Write(&b2))

	require.Equal(t, b1.Bytes(), b2.Bytes())
}

// Ensures that two fresh documents created from the same input yield identical bytes.
func TestDeterministicWrite_DifferentInstances(t *testing.T) {
	rd1 := buildElaborateDoc(t)
	rd2 := buildElaborateDoc(t)

	var b1, b2 bytes.Buffer
	require.NoError(t, rd1.Write(&b1))
	require.NoError(t, rd2.Write(&b2))

	require.Equal(t, b1.Bytes(), b2.Bytes())
}

// buildElaborateDoc creates a document that exercises multiple features
// (headings, styled runs, ordered/bulleted lists with nesting, and a table)
// to verify byte-for-byte determinism of the writer.
func buildElaborateDoc(t *testing.T) *docxpkg.RootDoc {
	t.Helper()

	rd, err := godocx.NewDocument()
	require.NoError(t, err)

	// Headings and styled paragraph content
	_, err = rd.AddHeading("Document Title", 0)
	require.NoError(t, err)

	p := rd.AddParagraph("A plain paragraph having some ")
	p.AddText("bold").Bold(true)
	p.AddText(" and some ")
	p.AddText("italic.").Italic(true)

	_, err = rd.AddHeading("Heading, level 1", 1)
	require.NoError(t, err)
	rd.AddParagraph("Intense quote").Style("Intense Quote")

	// Ordered list (abstract 1), with nesting
	ord := rd.NewListInstance(1)
	p = rd.AddParagraph("Ordered 1")
	p.Numbering(ord, 0)
	p = rd.AddParagraph("Ordered 2")
	p.Numbering(ord, 0)
	p = rd.AddParagraph("Nested 2.1")
	p.Numbering(ord, 1)
	p = rd.AddParagraph("Nested 2.2")
	p.Numbering(ord, 1)

	// Bulleted list (abstract 2)
	bul := rd.NewListInstance(2)
	p = rd.AddParagraph("Bullet 1")
	p.Numbering(bul, 0)
	p = rd.AddParagraph("Bullet 2")
	p.Numbering(bul, 0)
	p = rd.AddParagraph("Nested Bullet 2.1")
	p.Numbering(bul, 1)

	// Simple table
	tbl := rd.AddTable()
	tbl.Style("LightList-Accent4")
	hdr := tbl.AddRow()
	hdr.AddCell().AddParagraph("Qty")
	hdr.AddCell().AddParagraph("ID")
	hdr.AddCell().AddParagraph("Description")

	row := tbl.AddRow()
	row.AddCell().AddParagraph("5")
	row.AddCell().AddParagraph("A001")
	row.AddCell().AddParagraph("Laptop")

	row = tbl.AddRow()
	row.AddCell().AddParagraph("10")
	row.AddCell().AddParagraph("B202")
	row.AddCell().AddParagraph("Smartphone")

	// Page break at the end to exercise that path as well
	rd.AddPageBreak()

	return rd
}
