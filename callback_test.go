package lbd

import (
	"testing"
)

func TestConstructRawTransaction(t *testing.T) {
	b := []byte{123, 34, 104, 101, 105, 103, 104, 116, 34, 58, 49, 57, 51, 51, 52, 48, 48, 53, 44, 34, 116, 120, 104, 97, 115, 104, 34, 58, 34, 49, 67, 70, 55, 65, 66, 67, 65, 51, 48, 70, 51, 50, 67, 67, 51, 49, 69, 67, 57, 52, 54, 69, 50, 69, 56, 53, 70, 65, 70, 48, 49, 65, 67, 66, 68, 70, 48, 65, 56, 48, 54, 53, 50, 49, 67, 54, 57, 66, 69, 49, 69, 52, 69, 69, 55, 69, 49, 56, 49, 53, 67, 48, 56, 34, 44, 34, 99, 111, 100, 101, 115, 112, 97, 99, 101, 34, 58, 34, 34, 44, 34, 99, 111, 100, 101, 34, 58, 48, 44, 34, 105, 110, 100, 101, 120, 34, 58, 48, 44, 34, 100, 97, 116, 97, 34, 58, 34, 34, 44, 34, 108, 111, 103, 115, 34, 58, 91, 123, 34, 109, 115, 103, 73, 110, 100, 101, 120, 34, 58, 48, 44, 34, 108, 111, 103, 34, 58, 34, 34, 44, 34, 101, 118, 101, 110, 116, 115, 34, 58, 91, 123, 34, 116, 121, 112, 101, 34, 58, 34, 109, 101, 115, 115, 97, 103, 101, 34, 44, 34, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 34, 58, 91, 123, 34, 107, 101, 121, 34, 58, 34, 97, 99, 116, 105, 111, 110, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 115, 101, 110, 100, 34, 125, 44, 123, 34, 107, 101, 121, 34, 58, 34, 109, 111, 100, 117, 108, 101, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 99, 111, 105, 110, 34, 125, 44, 123, 34, 107, 101, 121, 34, 58, 34, 115, 101, 110, 100, 101, 114, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 116, 108, 105, 110, 107, 49, 97, 109, 56, 116, 112, 100, 110, 55, 122, 53, 53, 97, 100, 107, 115, 55, 115, 100, 121, 117, 115, 55, 99, 113, 112, 122, 56, 103, 102, 116, 117, 52, 103, 108, 114, 101, 112, 103, 34, 125, 93, 125, 44, 123, 34, 116, 121, 112, 101, 34, 58, 34, 116, 114, 97, 110, 115, 102, 101, 114, 34, 44, 34, 97, 116, 116, 114, 105, 98, 117, 116, 101, 115, 34, 58, 91, 123, 34, 107, 101, 121, 34, 58, 34, 115, 101, 110, 100, 101, 114, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 116, 108, 105, 110, 107, 49, 97, 109, 56, 116, 112, 100, 110, 55, 122, 53, 53, 97, 100, 107, 115, 55, 115, 100, 121, 117, 115, 55, 99, 113, 112, 122, 56, 103, 102, 116, 117, 52, 103, 108, 114, 101, 112, 103, 34, 125, 44, 123, 34, 107, 101, 121, 34, 58, 34, 114, 101, 99, 105, 112, 105, 101, 110, 116, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 116, 108, 105, 110, 107, 49, 113, 109, 118, 54, 107, 102, 119, 114, 119, 100, 97, 51, 115, 104, 52, 102, 114, 109, 57, 107, 120, 122, 52, 99, 104, 55, 120, 48, 100, 50, 56, 101, 107, 122, 120, 104, 107, 53, 34, 125, 44, 123, 34, 107, 101, 121, 34, 58, 34, 97, 109, 111, 117, 110, 116, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 49, 116, 99, 111, 110, 121, 34, 125, 93, 125, 93, 125, 93, 44, 34, 105, 110, 102, 111, 34, 58, 34, 34, 44, 34, 103, 97, 115, 87, 97, 110, 116, 101, 100, 34, 58, 54, 49, 48, 51, 50, 44, 34, 103, 97, 115, 85, 115, 101, 100, 34, 58, 52, 57, 53, 56, 52, 44, 34, 116, 120, 34, 58, 123, 34, 116, 121, 112, 101, 34, 58, 34, 99, 111, 115, 109, 111, 115, 45, 115, 100, 107, 47, 83, 116, 100, 84, 120, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 123, 34, 109, 115, 103, 34, 58, 91, 123, 34, 116, 121, 112, 101, 34, 58, 34, 99, 111, 105, 110, 47, 77, 115, 103, 83, 101, 110, 100, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 123, 34, 102, 114, 111, 109, 34, 58, 34, 116, 108, 105, 110, 107, 49, 97, 109, 56, 116, 112, 100, 110, 55, 122, 53, 53, 97, 100, 107, 115, 55, 115, 100, 121, 117, 115, 55, 99, 113, 112, 122, 56, 103, 102, 116, 117, 52, 103, 108, 114, 101, 112, 103, 34, 44, 34, 116, 111, 34, 58, 34, 116, 108, 105, 110, 107, 49, 113, 109, 118, 54, 107, 102, 119, 114, 119, 100, 97, 51, 115, 104, 52, 102, 114, 109, 57, 107, 120, 122, 52, 99, 104, 55, 120, 48, 100, 50, 56, 101, 107, 122, 120, 104, 107, 53, 34, 44, 34, 97, 109, 111, 117, 110, 116, 34, 58, 91, 123, 34, 100, 101, 110, 111, 109, 34, 58, 34, 116, 99, 111, 110, 121, 34, 44, 34, 97, 109, 111, 117, 110, 116, 34, 58, 49, 125, 93, 125, 125, 93, 44, 34, 102, 101, 101, 34, 58, 123, 34, 103, 97, 115, 34, 58, 54, 49, 48, 51, 50, 44, 34, 97, 109, 111, 117, 110, 116, 34, 58, 91, 93, 125, 44, 34, 109, 101, 109, 111, 34, 58, 34, 34, 44, 34, 115, 105, 103, 110, 97, 116, 117, 114, 101, 115, 34, 58, 91, 123, 34, 112, 117, 98, 75, 101, 121, 34, 58, 123, 34, 116, 121, 112, 101, 34, 58, 34, 116, 101, 110, 100, 101, 114, 109, 105, 110, 116, 47, 80, 117, 98, 75, 101, 121, 83, 101, 99, 112, 50, 53, 54, 107, 49, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 34, 65, 52, 78, 97, 65, 47, 120, 105, 66, 110, 77, 69, 70, 80, 108, 56, 75, 88, 84, 54, 76, 73, 113, 69, 115, 115, 98, 101, 89, 122, 82, 69, 103, 88, 80, 83, 122, 101, 51, 90, 119, 88, 66, 49, 34, 125, 44, 34, 115, 105, 103, 110, 97, 116, 117, 114, 101, 34, 58, 34, 73, 86, 67, 102, 48, 121, 52, 113, 78, 56, 73, 81, 79, 75, 69, 48, 52, 121, 67, 83, 109, 115, 84, 114, 78, 77, 73, 100, 79, 100, 76, 83, 108, 110, 53, 83, 55, 65, 78, 81, 66, 49, 112, 108, 104, 119, 110, 55, 68, 115, 71, 106, 50, 68, 112, 121, 108, 98, 116, 100, 106, 66, 54, 79, 114, 85, 76, 105, 71, 104, 119, 88, 97, 112, 117, 111, 52, 90, 107, 87, 97, 80, 116, 97, 48, 81, 61, 61, 34, 125, 93, 125, 125, 44, 34, 116, 105, 109, 101, 115, 116, 97, 109, 112, 34, 58, 49, 54, 52, 51, 49, 56, 56, 50, 56, 56, 48, 48, 48, 125}
	ret, err := ConstructRawTransaction(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
