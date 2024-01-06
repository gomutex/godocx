package main

import (
	"log"

	"github.com/gomutex/godocx"
)

func main() {
	// docx, err := godocx.NewDocument()
	docx, err := godocx.OpenDocument("./testdata/test.docx")
	if err != nil {
		log.Fatal(err)
	}
	// for _, rel := range docx.DocRelation.Relationships {
	// 	fmt.Println(rel.Type, rel.Target)
	// }
	// fmt.Println(docx)

	// fmt.Println(docx.CoreProperties)

	// fmt.Println(docx.FileMap)

	p := docx.AddParagraph()
	p.AddText("Hello, world!")

	// nextPara := docx.AddParagraph()
	// nextPara.AddLink("google", `http://google.com`)

	err = docx.SaveTo("./test.docx")
	if err != nil {
		log.Fatal(err)
	}
}
