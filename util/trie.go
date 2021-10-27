package util

type Node struct {
	Url      string
	Children map[string]*Node
}

func (n *Node) Insert(url string, parts []string, index int) {
	if len(parts) == index {
		n.Url = url
		return
	}
	part := parts[index]
	child := n.match(part)
	if child == nil {
		child = &Node{
			Children: make(map[string]*Node),
		}
		n.Children[part] = child
	}
	child.Insert(url, parts, index + 1)
}

func (n *Node) match(part string) *Node {
	return n.Children[part]
}

func (n *Node) Search(parts []string, index int) string {
	if index == len(parts) {
		return n.Url
	}
	part := parts[index]
	child := n.Children[part]
	if child == nil {
		for k, v := range n.Children {
			if k[0] == ':' {
				searchRes := v.Search(parts, index + 1)
				if searchRes != "" {
					return searchRes
				}
			}
		}
		return ""
	}
	return child.Search(parts, index + 1)
}
