// 路由匹配采用前缀树的形式

package biubiu

import "strings"

type TrieTree struct {
	root *node
}

func NewTrie() *TrieTree {
	return &TrieTree{root: &node{}}
}

func (t *TrieTree) Search(fullPath string) interface{} {
	return t.root.search(fullPath)
}

func (t *TrieTree) Insert(fullPath string, value interface{}) {
	t.root.insert(fullPath, value)
}

func (t *TrieTree) Clear() {
	t.root = &node{}
}

type node struct {
	fullPath string
	partPath string
	isWild   bool // 是否模糊匹配：*/:var
	value    interface{}
	children []*node
}

func parsePath(fullPath string) []string {
	splitPath := strings.Split(fullPath, "/")

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

func (n *node) insert(fullPath string, value interface{}) {
	parts := parsePath(fullPath)
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
	n.fullPath = fullPath
	n.value = value
}

func (n *node) search(fullPath string) interface{} {
	parts := parsePath(fullPath)
	var children = []*node{n}
	for _, part := range parts {
		children = n.doSearch(part, children)
	}
	if len(children) == 0 {
		return nil
	}
	return children[0].value
}

func (n *node) doSearch(part string, nodes []*node) []*node {
	var matched []*node
	for _, node := range nodes {
		matched = append(matched, node.matchChildren(part)...)
	}
	return matched
}
