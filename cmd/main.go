package main

import (
	"flag"
	"log"

	"github.com/gomutex/godocx"
)

func main() {

	inputFileFlag := flag.String("input", "./testdata/test.docx", "Path to the input file")
	outputFileFlag := flag.String("output", "/tmp/test.docx", "Path to the output file")
	flag.Parse()

	// docx, err := godocx.NewDocument()
	docx, err := godocx.OpenDocument(*inputFileFlag)
	if err != nil {
		log.Fatal(err)
	}
	// for _, rel := range docx.DocRelation.Relationships {
	// 	fmt.Println(rel.Type, rel.Target)
	// }
	// fmt.Println(docx)

	// fmt.Println(docx.CoreProperties)

	// fmt.Println(docx.FileMap)

	// p := docx.AddParagraph()
	// p.AddText("Hello, world!")

	// nextPara := docx.AddParagraph()
	// nextPara.AddLink("google", `http://google.com`)

	err = docx.SaveTo(*outputFileFlag)
	if err != nil {
		log.Fatal(err)
	}
}
