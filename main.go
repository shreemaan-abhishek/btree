package main

import (
	"fmt"
	"strings"
)

// todo:
// - handle equal values collision
// - handle balancing
//

type Node struct {
	left  *Node
	right *Node
	value *int
}

func tree_print(n *Node) string {
	if n == nil || n.value == nil {
		return ""
	}
	var b strings.Builder

	var walk func(n *Node, prefix string, isLeft bool, isRoot bool)
	walk = func(n *Node, prefix string, isLeft bool, isRoot bool) {
		if n == nil || n.value == nil {
			return
		}

		if n.right != nil && n.right.value != nil {
			newPrefix := prefix
			if isLeft && !isRoot {
				newPrefix += "|  "
			} else {
				newPrefix += "   "
			}
			walk(n.right, newPrefix, false, false)
		}

		b.WriteString(prefix)
		if len(prefix) > 0 {
			if isLeft {
				b.WriteString("`-- ")
			} else {
				b.WriteString("+-- ")
			}
		}
		b.WriteString(fmt.Sprintf("(%d)\n", *n.value))

		if n.left != nil && n.left.value != nil {
			newPrefix := prefix
			if isLeft {
				newPrefix += "   "
			} else {
				newPrefix += "|  "
			}
			walk(n.left, newPrefix, true, false)
		}
	}

	walk(n, "", true, true)
	return b.String()
}

var tree Node

func main() {
	for {
		fmt.Println("What do you want to insert into the tree?")
		var input int
		if _, err := fmt.Scanf("%d\n", &input); err != nil {
			fmt.Printf("failed to read, err: %s", err.Error())
		}
		tree = *insert(&tree, input)
		fmt.Println(tree_print(&tree))
	}

}

func insert(tree *Node, input int) *Node {
	if tree == nil {
		n := &Node{
			value: &input,
		}
		return n
	}
	if tree.value == nil {
		tree.value = &input
		return tree
	}
	value := *tree.value
	if input < value {
		tree.left = insert(tree.left, input)
		return tree
	}
	tree.right = insert(tree.right, input)
	return tree
}

func TreeDepth(node *Node) int {
	if node == nil {
		return 0
	}
	if node.left == nil {
		return 1
	}
	return 1 + TreeDepth(node.left)
}
