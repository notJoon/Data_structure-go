/*
Binary Search Tree

				100    <-- Root
			   /    \
			52      203  <- Parent
		   /  \     /  \
		  19  76   150  310  <- Children
		 /  \   \       /
		7   24  88     276   <- Leaves

 if children > parents -> children : right
 if children < parents -> children : left

 Time complexity
  - Balanced Tree : O(log n)
  - Unbalanced Tree : O(n)   -> worst case
*/
package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key : k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key : k}
		} else {
			n.Left.Insert(k)
		}
	}
}	

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}

/*
func main() {
	tree := &Node{Key : 100}
	fmt.Println(tree)

	tree.Insert(53)
	tree.Insert(413)
	tree.Insert(24)
	tree.Insert(150)
	tree.Insert(44)
	tree.Insert(23)
	tree.Insert(89)
	tree.Insert(245)
	tree.Insert(1)

	fmt.Println(tree.Search(1))
	print(count)
}
*/