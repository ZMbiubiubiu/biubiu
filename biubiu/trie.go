// 路由匹配采用前缀树的形式

package biubiu

import "strings"

type TrieTree struct {
	root *node
}

func NewTrie() *TrieTree {
	return &TrieTree{root: &node{}}
}

func (t *TrieTree) Search(pathPattern string) (HandlerFunc, string) {
	node := t.root.search(pathPattern)
	if node == nil {
		return nil, ""
	}
	return node.handler, node.pathPattern
}

func (t *TrieTree) Insert(pathPattern string, handler HandlerFunc) {
	t.root.insert(pathPattern, handler)
}

func (t *TrieTree) Clear() {
	t.root = &node{}
}

type node struct {
	pathPattern string
	partPath    string
	isWild      bool // 是否模糊匹配：*/:var
	handler     HandlerFunc
	children    []*node
}

func parsePath(pathPattern string) []string {
	splitPath := strings.Split(pathPattern, "/")

	var parts = make([]string, 0, len(splitPath))
	for _, v := range splitPath {
		if v == "" {
			continue
		}
		parts = append(parts, v)
		if v[0] == '*' {
			break
		}
	}
	return parts
}

func (n *node) matchFirstChild(part string) *node {
	for _, child := range n.children {
		if child.isWild || child.partPath == part {
			return child
		}
	}
	return nil
}
func (n *node) matchChildren(part string) []*node {
	var children = make([]*node, 0, len(n.children))
	for _, child := range n.children {
		if child.isWild || child.partPath == part {
			children = append(children, child)
		}
	}
	return children
}

func (n *node) insert(pathPattern string, handler HandlerFunc) {
	parts := parsePath(pathPattern)
	for _, part := range parts {
		child := n.matchFirstChild(part)
		if child == nil {
			child = &node{
				partPath: part,
				isWild:   part[0] == '*' || part[0] == ':',
			}
			n.children = append(n.children, child)
		}
		n = child
	}
	// for the last node
	n.pathPattern = pathPattern
	n.handler = handler
}

func (n *node) search(pathPattern string) *node {
	parts := parsePath(pathPattern)
	var children = []*node{n}
	for _, part := range parts {
		children = n.doSearch(part, children)
	}
	if len(children) == 0 {
		return nil
	}
	return children[0]
}

func (n *node) doSearch(part string, nodes []*node) []*node {
	var matched []*node
	for _, node := range nodes {
		matched = append(matched, node.matchChildren(part)...)
	}
	return matched
}
