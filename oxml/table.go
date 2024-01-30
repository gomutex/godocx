package oxml

import "github.com/gomutex/godocx/oxml/elements"

func (rd *RootDoc) AddTable() *elements.Table {
	tbl := elements.DefaultTable()
	tbl.Rows = []elements.TableChild{}
	tbl.Grid = elements.DefaultTableGrid()

	rd.Document.Body.Children = append(rd.Document.Body.Children, DocumentChild{
		Table: tbl,
	})

	return tbl
}
