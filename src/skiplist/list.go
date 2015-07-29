package skiplist

import (
	"fmt"
	"math/rand"
)

// A SkipList is a map-like data structure that maintains an ordered
// collection of key-value pairs. Insertion, lookup, and deletion are
// all O(log n) operations. A SkipList can efficiently store up to
// 2^MaxLevel items.
//
// sorted as increasment
//
type SkipList struct {
	head     *Node // first node of skip list
	end      *Node // last node of skip list
	length   int64 // the length of skip list
	maxLevel int   // max level num
}

// New skip list
func NewSkipList(maxLevel int) *SkipList {
	obj := new(SkipList)
	obj.maxLevel = maxLevel
	obj.head = NewNode(maxLevel, maxLevel)
	return obj
}

// Get length of skip list
func (this *SkipList) Length() int64 {
	return this.length
}

// Get random level
func (this *SkipList) randomLevel() int {
	n := 0
	for ; n < this.getMaxLevel() && rand.Float64() < 0.25; n++ {
	}
	return n
}

// Set max level
func (this *SkipList) SetMaxLevel(maxLevel int) {
	this.maxLevel = maxLevel
}

// Max level
func (this *SkipList) getMaxLevel() int {
	return this.maxLevel
}

// print list
func (this *SkipList) Print(callback func(data Data)) {
	for i := this.getMaxLevel() - 1; i >= 0; i-- {
		node := this.head
		fmt.Print(fmt.Sprintf("level%d:", i))
		fmt.Print("head->")
		for len(node.next) > 0 && node.next[i] != nil {
			node = node.next[i]
			fmt.Print(fmt.Sprintf("(%d)", node.width[i]))
			callback(node.data)
			fmt.Print("->")
		}
		fmt.Print("nil\n")
	}
}

// get left nodes, left nodes is level0-levelmax nodes every in which is
// less than @data
func (this *SkipList) getLNodes(data Data, lNodes []*Node, lWidths []int64) *Node {
	var find *Node
	node := this.head
	width := int64(0)
	for i := this.getMaxLevel() - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].data.Less(data) {
			node = node.next[i]
			width += node.width[i]
			if node.data.Equal(data) {
				find = node
			}
		}
		if lNodes != nil {
			lNodes[i] = node
		}
		if lWidths != nil {
			lWidths[i] = width
		}
	}
	// fmt.Println(fmt.Sprintf("data %v, width %d", data, width))
	if find != nil {
		return find
	}
	if node.next != nil && len(node.next) > 0 {
		return node.next[0]
	}
	return nil
}

// Get node from skip list
// by the data of node
func (this *SkipList) Get(data Data) *Node {
	return this.getLNodes(data, nil, nil)
}

// del old node start by @lNodes
func (this *SkipList) del(oldNode *Node, lNodes []*Node) {
	if this.end == oldNode {
		this.end = oldNode.pre
	}
	for i := len(lNodes) - 1; i >= 0; i-- {
		lNode := lNodes[i]
		if lNode.next[i] == nil || len(lNode.next) <= 0 {
			continue
		}
		if lNode.next[i].data.Equal(oldNode.data) {
			if lNode.next[i].next[i] == nil || len(lNode.next[i].next) <= 0 {
				lNode.next[i] = nil
				continue
			}
			if i == 0 {
				lNode.next[i].next[i].pre = lNode.next[i]
			}
			lNode.next[i].next[i].width[i] += (lNode.next[i].width[i] - 1)
			lNode.next[i] = lNode.next[i].next[i]
		} else {
			lNode.next[i].width[i]--
		}
	}
	this.length--
	fmt.Println("delete success")
}

// delete an old data if it exists
func (this *SkipList) Del(old Data) {
	maxLevel := this.getMaxLevel()

	//find node that node.next.data >= data for each level
	lNodes := make([]*Node, maxLevel, maxLevel)
	lWidths := make([]int64, maxLevel, maxLevel)

	//find exist node if found then delete it first
	oldNode := this.getLNodes(old, lNodes, lWidths)
	if oldNode != nil && oldNode.data.Equal(old) {
		this.del(oldNode, lNodes)
	}
}

// @data is the new data or will update data into the skip list
// if update a new data which equal the old one but the value (which
// used to less compare is not similar to the old one)
// then you should record the old one and input it into Set
// by @old
func (this *SkipList) Set(data Data, old Data) {
	//find exist node if found then delete it first
	if old != nil {
		this.Del(old)
	}

	maxLevel := this.getMaxLevel()

	//find node that node.next.data >= data for each level
	lNodes := make([]*Node, maxLevel, maxLevel)
	lWidths := make([]int64, maxLevel, maxLevel)
	this.getLNodes(data, lNodes, lWidths)

	level := this.randomLevel() + 1
	newNode := NewNode(level, this.getMaxLevel())
	newNode.data = data
	for i := level; i < this.getMaxLevel(); i++ {
		if lNodes[i].next[i] != nil {
			lNodes[i].next[i].width[i]++
		}
	}
	for i := 0; i < level; i++ {
		if lNodes[i].next != nil {
			newNode.next[i] = lNodes[i].next[i]
		} else {
			newNode.next[i] = nil
		}
		lNodes[i].next[i] = newNode

		if i == 0 {
			newNode.width[i] = 1
		} else {
			newNode.width[i] = lWidths[0] + 1 - lWidths[i]
		}
		if newNode.next[i] != nil {
			newNode.next[i].width[i] = newNode.next[i].width[i] - newNode.width[i] + 1
		}
	}

	newNode.pre = lNodes[0]
	if newNode.next[0] != nil {
		newNode.next[0].pre = newNode
	}

	this.length++

	if this.end == nil || this.end.data.Less(newNode.data) {
		this.end = newNode
	}
}
