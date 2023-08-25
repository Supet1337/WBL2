package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	testingStrings := []struct {
		inputStr    string
		expectedStr string
		err         error
	}{
		{
			"a4bc2d5e",
			"aaaabccddddde",
			nil,
		},
		{
			"abcd",
			"abcd",
			nil,
		},
		{
			"45",
			"",
			errors.New("Incorrect"),
		},
	}
	for _, testItem := range testingStrings {
		s, err := Unpack(testItem.inputStr)
		if s != testItem.expectedStr {
			t.Errorf("fail test with string: %v", testItem.inputStr)
		}
		if err != nil && s != testItem.expectedStr {
			t.Errorf("fail test with string: %v err: %v", testItem.inputStr, testItem.err)
		}
	}
}
