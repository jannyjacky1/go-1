package hw_2

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {

	testStrings := []string{"a", "a5", "5a", "m24y40string90"}
	testStringsExpectedResults := []string{"a", "aaaaa", "", "mmmmmmmmmmmmmmmmmmmmmmmmyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyystringgggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"}
	testStringsExpectedErrors := []error{nil, nil, errors.New(""), nil}

	for i, str := range testStrings {
		resStr, resErr := UnpackString(str)

		if resStr != testStringsExpectedResults[i] {
			t.Fatalf("bad unpack for %s: got %s expected %s", str, resStr, testStringsExpectedResults[i])
		}
		if resErr != testStringsExpectedErrors[i] {
			t.Fatalf("bad unpack for %s: got error %v expected %v", str, resErr, testStringsExpectedErrors[i])
		}
	}

}
