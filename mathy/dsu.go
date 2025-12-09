package mathy

type node struct {
	value  any
	parent *node
	size   int
}

type DSU struct {
	nodes map[any]*node
}

func NewDSU() *DSU {
	return &DSU{map[any]*node{}}
}

func (d DSU) Contains(x any) bool {
	_, ok := d.nodes[x]
	return ok
}

func (d *DSU) Add(x any) {
	if !d.Contains(x) {
		d.nodes[x] = &node{value: x, parent: nil, size: 1}
	}
}

func (d *DSU) Find(x any) any {
	if !d.Contains(x) {
		return nil
	}

	node := d.nodes[x]
	root := node

	for root.parent != nil {
		root = root.parent
	}

	for node.parent != nil {
		node.parent, node = root, node.parent
	}

	return root.value
}

// Returns the size of the sub-graph with x and y.
func (d *DSU) Union(x, y any) int {
	rootX := d.Find(x)
	rootY := d.Find(y)

	if rootX == nil || rootY == nil || rootX == rootY {
		return 0
	}

	nodeX := d.nodes[rootX]
	nodeY := d.nodes[rootY]

	if nodeX.size <= nodeY.size {
		nodeX, nodeY = nodeY, nodeX
	}

	nodeY.parent = nodeX
	nodeX.size += nodeY.size

	return nodeX.size
}
