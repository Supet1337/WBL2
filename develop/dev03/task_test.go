package test

import (
	"dev03/internal"
	"reflect"
	"testing"
)

func TestSortByColumn(t *testing.T) {
	test := []struct {
		input  []string
		output []string
		flag   *internal.Flag
	}{
		{
			[]string{"bbb 4", "bbb 3", "aaa 2", "ccc 1"},
			[]string{"ccc 1", "aaa 2", "bbb 3", "bbb 4"},
			&internal.Flag{K: 2},
		},
	}
	for _, testData := range test {
		if !reflect.DeepEqual(internal.SortByColumn(testData.input, testData.flag), testData.output) {
			t.Errorf("Column sort error:\n%v\n%v\n", testData.output, internal.SortByColumn(testData.input, testData.flag))
		}
	}
}
