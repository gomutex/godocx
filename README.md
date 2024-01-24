# Godocx

Godocx is a Golang library for creating and modifying DOCX (Microsoft Word) documents.

## Usage
Here's a simple example of how you can use Godocx to create and modify DOCX documents:

```go
package main

import (
	"log"

	"github.com/gomutex/godocx"
)

func main() {
	// Open an existing DOCX document
	docx, err := godocx.OpenDocument("./testdata/test.docx")
	if err != nil {
		log.Fatal(err)
	}

	// Add a new paragraph to the document
	_ = docx.AddParagraph("Hello World")
	p := docx.AddEmptyParagraph()
	_ = p.AddText("Add Paragraph and get `Run` instance")

	// Save the modified document to a new file
	err = docx.SaveTo("./test_modified.docx")
	if err != nil {
		log.Fatal(err)
	}
}

```

## Inspiration
This GoDocx Library draws inspiration from two renowned libraries in the programming world - python-docx and docx-rs (Rust). 


