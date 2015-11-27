package kolonsecfg

import ()

const (
	ONCE_CAP_LEN = 10
)

type Nodes struct {
	Ns []*Node
	l  int
	c  int
}

func (ns *Nodes) Add(n *Node) {
	if ns.l >= ns.c {
		// 需要进行扩充容器
		cn := make([]*Node, ONCE_CAP_LEN)
		ns.Ns = append(ns.Ns, cn...)
		ns.c += ONCE_CAP_LEN
	}

	ns.Ns[ns.l] = n
	ns.l++
}

func (ns *Nodes) Get() []*Node {
	return ns.Ns[:ns.l]
}

func NewNodes() *Nodes {
	return &Nodes{
		Ns: make([]*Node, ONCE_CAP_LEN),
		c:  ONCE_CAP_LEN,
		l:  0,
	}
}

type Node struct {
	key string
	*Value
	farther *Node
	childs  map[string]*Nodes
}

func (n *Node) SetKey(key string) {
	n.key = key
}

func (n *Node) SetValue(value *Value) {
	n.Value = value
}

func (n *Node) SetFarther(farther *Node) {
	n.farther = farther
}

func (n *Node) AddChild(child *Node) *Node {
	nodes, ok := n.childs[child.key]
	if !ok { // 如果不存在该 key 那么就创建
		nodes = NewNodes()
		n.childs[child.key] = nodes
	}
	// 将该节点值添加到节点映射中
	nodes.Add(child)
	child.SetFarther(n)
	return n
}

func (n *Node) Childs(key string) []*Node {
	nodes, ok := n.childs[key]
	if !ok {
		return nil
	}
	return nodes.Get()
}

func (n *Node) Child(key string) *Node {
	nodes, ok := n.childs[key]
	if !ok {
		return nil
	}
	return nodes.Get()[0]
}

func (n *Node) Dump(suffix string) string {
	ret := ""
	// node 值不为空时进行打印
	if n.key != "" && n.Value != nil && n.Value.ValueType != INVALID {
		ret += suffix + n.key + " " + n.Value.GetString() + "\n"
	} else if n.key != "" {
		ret += suffix + n.key + " {\n"
		for _, nodes := range n.childs {
			for _, node := range nodes.Get() {
				ret += node.Dump(suffix + "\t")
			}
		}
		ret += suffix + "}\n"
	} else {
		for _, nodes := range n.childs {
			for _, node := range nodes.Get() {
				ret += node.Dump(suffix + "\t")
			}
		}
	}
	return ret
}

func (n *Node) AddChilds(childs []*Node) {
	//	n.child = append(n.child, childs...)
	//	for _, i := range childs {
	//		childs[i].SetFarther(n)
	//	}
}

func NewNode(key string) *Node {
	return &Node{
		key:    key,
		childs: make(map[string]*Nodes),
	}
}
