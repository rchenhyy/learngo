package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

func newTreeNode(value int) *treeNode { // 工厂函数一般返回一个struct的地址
	return &treeNode{value: value}
}

func (node *treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignore!")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // new vs make
	root.left.right = newTreeNode(2)

	nodes := []treeNode{
		{value: 3}, // type name omitted
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

	root.traverse()

	var t treeNode
	fmt.Println(t)	//
}
