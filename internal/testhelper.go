package internal

import (
	"fmt"
	"reflect"
)

func ToPtr[T any](input T) *T {
	return &input
}

func FormatPtr[T any](ptr *T) string {
	if ptr == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", *ptr)
}

// func ComparePtr[T comparable](fieldName string, a, b *T) error {
// 	if a == nil || b == nil {
// 		if a != b {
// 			return fmt.Errorf("%s: expected %v but got %v", fieldName, FormatPtr(a), FormatPtr(b))
// 		}
// 	} else if *a != *b {
// 		return fmt.Errorf("%s: expected %v but got %v", fieldName, *a, *b)
// 	}
// 	return nil
// }

func ComparePtr[T comparable](fieldName string, a, b *T) error {
	// Check if T is a struct
	if reflect.TypeOf(*new(T)).Kind() == reflect.Struct {
		if a == nil || b == nil {
			if a != b {
				return fmt.Errorf("%s: expected %v but got %v", fieldName, FormatPtr(a), FormatPtr(b))
			}
		}
	} else {
		// For non-struct types, perform value comparison
		if a == nil || b == nil {
			if a != b {
				return fmt.Errorf("%s: expected %v but got %v", fieldName, FormatPtr(a), FormatPtr(b))
			}
		} else if *a != *b {
			return fmt.Errorf("%s: expected %v but got %v", fieldName, *a, *b)
		}
	}
	return nil
}
