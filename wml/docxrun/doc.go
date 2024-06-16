// Package docxrun provides functionalities for managing runs within DOCX (WordprocessingML) documents in Go.
//
// # Run
//
// A run in a DOCX document defines a region of text with a common set of properties, represented by the <w:r> element
// (§2.3.2.23) in WordprocessingML. It allows specifying formatting properties, such as bold and italic, uniformly
// applying them to all content within the run.
//
// For more details on WordprocessingML runs, refer to:
// http://webapp.docx4java.org/OnlineDemo/ecma376/WordML/Run_1.html
//
// Package docxrun provides methods to manipulate and control runs, enabling users to handle text formatting and
// manage properties within DOCX files.
package docxrun
