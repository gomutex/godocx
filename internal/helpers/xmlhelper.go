package helpers

import "encoding/xml"

// SkipUntilEnd is a helper function that reads an XML stream and skips over
// any empty tags until it reaches the specified end tag.
func SkipUntilEnd(d *xml.Decoder, end xml.EndElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := token.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if elem == end {
				return nil
			}
		}
	}
}
