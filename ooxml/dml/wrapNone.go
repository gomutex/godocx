package dml

// No Text Wrapping
type WrapNone struct{}

// func (w *WrapNone) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	start.Name.Local = "wp:wrapNone"

// 	err := e.EncodeToken(start)
// 	if err != nil {
// 		return err
// 	}
// 	return e.EncodeToken(xml.EndElement{Name: start.Name})
// }
