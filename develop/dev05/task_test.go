package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	test := []struct {
		input  TextFile
		output []string
	}{
		{
			input: TextFile{
				Flags: &Flags{v: true},
				dateStr: []string{"aa",
					"v",
					"g",
					"gg",
					"uuu",
				},
				path: "in.txt",
				word: "a",
			},
			output: []string{
				"v",
				"g",
				"gg",
				"uuu",
			},
		},
	}
	for _, testData := range test {
		inpt, _ := grep(&testData.input)
		if !reflect.DeepEqual(inpt, testData.output) {
			t.Errorf("grep fail:\nExpected: %v\n Recieve %v\n", testData.output, inpt)
		}
	}
}
