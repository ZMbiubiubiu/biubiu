// 路由匹配采用前缀树的形式

package biubiu

import "strings"

type node struct {
	fullPath string
	partPath string
	isWild   bool // 是否模糊匹配：*/:var
	handler  interface{}
	children []*node
}

func NewNode() *node {
	return &node{}
}

func (n *node) parsePath(fullPath string) []string {
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
