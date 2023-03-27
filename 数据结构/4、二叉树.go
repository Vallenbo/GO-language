package main

import "fmt"

type TreeNode struct { // 定义二叉树结构体
	data  int       // 数据域
	left  *TreeNode // 左子树指针
	right *TreeNode // 右子树指针
}

func CreateBinaryTree(data []int) *TreeNode { // 创建二叉树
	if len(data) == 0 {
		return nil
	}
	// 取出中间位置的元素作为根节点
	mid := len(data) / 2
	root := &TreeNode{data: data[mid]}
	// 递归地创建左子树和右子树
	root.left = CreateBinaryTree(data[:mid])
	root.right = CreateBinaryTree(data[mid+1:])
	return root
}

func PreOrderTraversal(root *TreeNode) { // 前序遍历
	if root == nil {
		return
	}

	fmt.Println(root.data) // 先访问根节点，再遍历左子树和右子树
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func InOrderTraversal(root *TreeNode) { // 中序遍历
	if root == nil {
		return
	}

	InOrderTraversal(root.left) // 先遍历左子树，再访问根节点，最后遍历右子树
	fmt.Println(root.data)
	InOrderTraversal(root.right)
}

func PostOrderTraversal(root *TreeNode) { // 后序遍历
	if root == nil {
		return
	}

	PostOrderTraversal(root.left) // 先遍历左子树和右子树，最后访问根节点
	PostOrderTraversal(root.right)
	fmt.Println(root.data)
}
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := CreateBinaryTree(data) // 创建二叉树
	PreOrderTraversal(root)        // 前序遍历
	InOrderTraversal(root)         // 中序遍历
	PostOrderTraversal(root)       // 后序遍历
}
