package docx

import (
	"testing"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
	"github.com/stretchr/testify/assert"
)

func assertParaText(t *testing.T, para *Paragraph, expected string) {
	t.Helper()

	assert.NotNil(t, para)
	ct := para.GetCT()
	assert.NotNil(t, ct)
	assert.GreaterOrEqual(t, len(ct.Children), 1)
	run := ct.Children[0].Run
	assert.NotNil(t, run)
	assert.GreaterOrEqual(t, len(run.Children), 1)
	text := run.Children[0].Text
	assert.NotNil(t, text)

	assert.Equal(t, text.Text, expected)

}

func TestAddParagraph(t *testing.T) {
	rd := setupRootDoc(t)
	para := rd.AddParagraph("Test paragraph")
	assertParaText(t, para, "Test paragraph")
}

func TestEmptyParagraph(t *testing.T) {
	rd := setupRootDoc(t)
	para := rd.AddEmptyParagraph()
	para.AddText("Test paragraph")
	assertParaText(t, para, "Test paragraph")
}

func TestParagraph_Style(t *testing.T) {
	f := func(styleValue string, expectedStyleValue string) {
		t.Helper()

		p := &Paragraph{}

		p.Style(styleValue)

		assert.NotNil(t, p.ct.Property)
		assert.NotNil(t, p.ct.Property.Style)
		assert.Equal(t, p.ct.Property.Style.Val, expectedStyleValue)
	}

	f("Heading1", "Heading1")
	f("Normal", "Normal")
}
func TestParagraph_Justification(t *testing.T) {
	f := func(justificationValue, expectedJustificationValue stypes.Justification) {
		t.Helper()

		p := &Paragraph{}

		p.Justification(justificationValue)

		assert.NotNil(t, p.ct.Property, "Expected ct.Property to be non-nil")
		assert.NotNil(t, p.ct.Property.Justification, "Expected ct.Property.Justification to be non-nil")
		assert.Equal(t, p.ct.Property.Justification.Val, expectedJustificationValue, "Paragraph.Justification() value mismatch")
	}

	f(stypes.JustificationCenter, stypes.JustificationCenter)
	f(stypes.JustificationLeft, stypes.JustificationLeft)
	f(stypes.JustificationRight, stypes.JustificationRight)
	f(stypes.JustificationBoth, stypes.JustificationBoth)
}

func TestParagraph_Numbering(t *testing.T) {
	f := func(id int, level int, expectedNumID int, expectedILvl int) {
		t.Helper()

		p := &Paragraph{}

		p.Numbering(id, level)

		assert.NotNil(t, p.ct.Property, "Expected ct.Property to be non-nil")
		assert.NotNil(t, p.ct.Property.NumProp, "Expected ct.Property.NumProp to be non-nil")
		assert.Equal(t, expectedNumID, p.ct.Property.NumProp.NumID.Val, "Paragraph.Numbering() NumID value mismatch")
		assert.Equal(t, expectedILvl, p.ct.Property.NumProp.ILvl.Val, "Paragraph.Numbering() ILvl value mismatch")
	}

	f(1, 0, 1, 0)
	f(2, 1, 2, 1)
	f(3, 2, 3, 2)
	f(4, 3, 4, 3)
}

func TestParagraph_Indent(t *testing.T) {
	f := func(indentValue, expectedIndentValue ctypes.Indent) {
		t.Helper()

		p := &Paragraph{}

		p.Indent(&indentValue)

		assert.NotNil(t, p.ct.Property, "Expected ct.Property to be non-nil")
		assert.NotNil(t, p.ct.Property.Indent, "Expected ct.Property.Indent to be non-nil")
		assert.Equal(t, p.ct.Property.Indent, &expectedIndentValue, "Paragraph.Indent() value mismatch")
	}

	var size6 int = 6
	var size360 int = 360
	var sizeu420 uint64 = 420

	indentLeft := ctypes.Indent{Left: &size360, Hanging: &sizeu420}
	indentRight := ctypes.Indent{Right: &size360, Hanging: &sizeu420}
	indentFirst := ctypes.Indent{FirstLine: &sizeu420}
	indentLeftChars := ctypes.Indent{LeftChars: &size6}
	indentRightChars := ctypes.Indent{RightChars: &size6}
	indentFirstChars := ctypes.Indent{FirstLineChars: &size6}

	f(indentLeft, indentLeft)
	f(indentRight, indentRight)
	f(indentFirst, indentFirst)
	f(indentLeftChars, indentLeftChars)
	f(indentRightChars, indentRightChars)
	f(indentFirstChars, indentFirstChars)
}

func TestParagraph_AddText(t *testing.T) {
	f := func(text string, expectedText string) {
		t.Helper()

		p := &Paragraph{
			ct: ctypes.Paragraph{
				Children: []ctypes.ParagraphChild{},
			},
		}

		run := p.AddText(text)

		assert.NotNil(t, run, "Expected AddText() to return a non-nil Run")
		assert.Equal(t, len(p.ct.Children), 1, "Expected one Run to be added to Paragraph")

		assert.NotNil(t, p.ct.Children[0].Run, "Expected ct.Children[0].Run to be non-nil")
		assert.GreaterOrEqual(t, len(p.ct.Children[0].Run.Children), 1, "Expected at least one Text element in Run")
		assert.NotNil(t, p.ct.Children[0].Run.Children[0].Text, "Expected Text element in Run to be non-nil")
		assert.Equal(t, p.ct.Children[0].Run.Children[0].Text.Text, expectedText, "Paragraph.AddText() Text value mismatch")
	}

	f("Hello, World!", "Hello, World!")
	f("Another test", "Another test")
	f("A third text", "A third text")
	f("Sample text", "Sample text")
}

func TestParagraph_AddRun(t *testing.T) {
	p := &Paragraph{
		ct: ctypes.Paragraph{
			Children: []ctypes.ParagraphChild{},
		},
	}

	run := p.AddRun()

	assert.NotNil(t, run, "Expected AddRun() to return a non-nil Run")

	assert.Equal(t, 1, len(p.ct.Children), "Expected one Run to be added to Paragraph")
	assert.NotNil(t, p.ct.Children[0].Run, "Expected ct.Children[0].Run to be non-nil")
	assert.Equal(t, run.ct, p.ct.Children[0].Run, "Expected the Run returned by AddRun() to match the Run added to Paragraph")

	assert.Empty(t, run.ct.Children, "Expected new Run to have no initial Children")

	assert.Equal(t, 0, len(p.ct.Children[0].Run.Children), "Expected the new Run to have no initial Children")
}
