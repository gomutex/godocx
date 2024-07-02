package docx

import (
	"log"
	"testing"

	"github.com/gomutex/godocx/wml/ctypes"
)

func setupRootDoc(t *testing.T) *RootDoc {
	log.Println("setting root doc")

	return &RootDoc{
		Path:        "/tmp/test",
		RootRels:    Relationships{},
		ContentType: ContentTypes{},
		Document: &Document{
			Body: &Body{},
		},
		DocStyles:  &ctypes.Styles{},
		rID:        1,
		ImageCount: 1,
	}
}
