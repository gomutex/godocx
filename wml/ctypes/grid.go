package ctypes

import (
	"encoding/xml"
	"fmt"
)

// Table Grid
type Grid struct {
	//1. Grid Column Definition
	Col []Column `xml:"gridCol,omitempty"`

	//2.Revision Information for Table Grid Column Definitions
	GridChange *GridChange `xml:"tblGridChange,omitempty"`
}

func (g Grid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:tblGrid"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	//1. Grid Column Definition
	for _, col := range g.Col {
		if err := col.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("table grid marshalling column: %w", err)
		}
	}

	//2.Revision Information for Table Grid Column Definitions
	if g.GridChange != nil {
		if err := g.GridChange.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("table grid marshalling gridchange : %w", err)
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
