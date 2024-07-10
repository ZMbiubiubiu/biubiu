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

	for _, table := range tables {
		got := parsePath(table.input)
		assert.Equal(t, len(table.output), len(got))
		for i, part := range got {
			assert.Equal(t, table.output[i], part)
		}
	}
}

func TestInsertAndSearch(t *testing.T) {
	var tables = []struct {
		fullPath string
		value    int
	}{
		{
			fullPath: "/hello/world",
			value:    2,
		},
		{
			fullPath: "/hello/",
			value:    1,
		},
	}

	trie := NewTrie()

	for _, table := range tables {
		trie.Insert(table.fullPath, table.value)
		got := trie.Search(table.fullPath)
		assert.NotNil(t, got)
		assert.Equal(t, table.value, got.(int))
		trie.Clear()
		got = trie.Search(table.fullPath)
		assert.Nil(t, got)
	}
}
