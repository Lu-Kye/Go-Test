package skiplist

// Data is the specific thing saved in
// the node
// data should implement less than function
// and equal function to compare with others
type Data interface {
	Less(Data) bool
	Equal(Data) bool
}

// A node is a container for key-value pairs
// that are stored in a skip list
type Node struct {
	next  []*Node // [0 : level0 next node, 1 : level1 next node]
	pre   *Node   // pre node
	width []int64 // [1 : distance to pre level0 node]
	data  Data    // real compare data
}

// New node
func NewNode(level int, maxLevel int) *Node {
	obj := &Node{}
	obj.next = make([]*Node, level, maxLevel)
	obj.width = make([]int64, level, maxLevel)
	return obj
}
