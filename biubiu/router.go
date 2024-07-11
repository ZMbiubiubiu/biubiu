package biubiu

import (
	"log"
	"strings"
)

type router struct {
	handlers map[string]*TrieTree
}

func newRouter() *router {
	return &router{handlers: make(map[string]*TrieTree)}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	_, ok := r.handlers[method]
	if !ok {
		r.handlers[method] = NewTrie()
	}
	r.handlers[method].Insert(pattern, handler)
}

func (r *router) getHandler(c *Context) (HandlerFunc, string) {
	trie, ok := r.handlers[c.Method]
	if !ok {
		return nil, ""
	}
	return trie.Search(c.Path)
}

func (r *router) handle(c *Context) {
	log.Printf("receive request || method=%s||path=%s", c.Method, c.Path)
	handler, pattern := r.getHandler(c)
	if handler == nil {
		c.NotFound("404 NOT FOUND: " + c.Path)
		return

	}
	setParams(c, pattern)
	handler(c)
}
func setParams(c *Context, pattern string) {
	params := make(map[string]string)
	patternParts := parsePath(pattern)
	pathParts := parsePath(c.Path)
	for index, patternPart := range patternParts {
		if len(patternPart) == 0 {
			continue
		}
		if patternPart[0] == ':' {
			params[patternPart[1:]] = pathParts[index]
		}
		if patternPart[0] == '*' && len(patternPart) > 1 {
			params[patternPart[1:]] = strings.Join(pathParts[index:], "/")
			break
		}
	}
	c.Params = params
}
