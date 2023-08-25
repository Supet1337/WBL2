package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestF(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output *map[string][]string
	}{
		{
			name:  "ok",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			output: &map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			response := Anagram(&item.input)
			assert.Equal(t, item.output, response)
		})
	}
}
