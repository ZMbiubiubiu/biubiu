package biubiu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePath(t *testing.T) {
	var tables = []struct {
		input  string
		output []string
	}{
		{
			input:  "/",
			output: []string{},
		},
		{
			input:  "/hello/world",
			output: []string{"hello", "world"},
		},
		{
			input:  "/hello//world",
			output: []string{"hello", "world"},
		},
		{
			input:  "/hello/:lang/bs",
			output: []string{"hello", ":lang", "bs"},
		},
		{
			input:  "/hello/*/world",
			output: []string{"hello", "*"},
		},
	}

	node := NewNode()

	for _, table := range tables {
		got := node.parsePath(table.input)
		assert.Equal(t, len(table.output), len(got))
		for i, part := range got {
			assert.Equal(t, table.output[i], part)
		}
	}
}
