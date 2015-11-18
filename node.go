package cfg

type Node struct {
	key     string
	value   *Value
	farther *Node
	child   []*Node
}

func (n *Node) SetFarther(farther *Node) {
	n.farther = farther
}

func (n *Node) AddChild(child *Node) {
	n.child = append(n.child, child)
	child.SetFarther(n)
}

func (n *Node) AddChilds(childs []*Node) {
	n.child = append(n.child, childs...)
	for _, i := range childs {
		childs[i].SetFarther(n)
	}
}

func NewNode() {
	return Node{}
}
