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
		handler  HandlerFunc
	}{
		{
			fullPath: "/hello/world",
			handler: HandlerFunc(func(c *Context) {
			}),
		},
		{
			fullPath: "/hello/",
			handler: HandlerFunc(func(c *Context) {
			}),
		},
	}

	trie := NewTrie()

	for _, table := range tables {
		trie.Insert(table.fullPath, table.handler)
		got := trie.Search(table.fullPath)
		assert.NotNil(t, got)
		trie.Clear()
		got = trie.Search(table.fullPath)
		assert.Nil(t, got)
	}
}
