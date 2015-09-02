package worldTree

import (
)

type TreeNode struct {
	Value int
	Lchild *TreeNode
	Rchild *TreeNode
}

func (this *TreeNode) WorldTreeInit() {
	this.Value = 1
	this.Lchild = nil
	this.Rchild = nil
}

func (this *TreeNode) GetNodeValue() (int) {
	return this.Value
}

func (this *TreeNode) InsertNode(parent *TreeNode, new *TreeNode) {
	if parent.Lchild == nil {
		parent.Lchild = new
	} else {
		parent = parent.Lchild
		for{
			if(parent.Rchild == nil) {
				parent.Rchild = new
				break
			} else {
				parent = parent.Rchild
			}
		}
	}
}

func (this *Tree) DeleteNode(parent *TreeNode, old *TreeNode) {	
}
