package DataStructure

import (
	"testing"
)

type TreeNode1 struct {
	Data  string     // 节点用来存放数据
	Left  *TreeNode1 // 左子树
	Right *TreeNode1 // 右字树
}

// a b c d e f
func Test_treeQueue(t *testing.T) { // 队列层序遍历打印深度
	root := &TreeNode1{Data: "A"}
	root.Left = &TreeNode1{Data: "B"}
	root.Right = &TreeNode1{Data: "C"}
	root.Left.Left = &TreeNode1{Data: "D"}
	root.Left.Right = &TreeNode1{Data: "E"}
	root.Right.Left = &TreeNode1{Data: "F"}

	queue := []*TreeNode1{root}
	var length int
	for len(queue) > 0 {
		length++
		curLayersNum := len(queue)
		for i := 0; i < curLayersNum; i++ {
			element := queue[0]
			println(element.Data)
			queue = queue[1:]
			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}
		}
	}
	println(length)
}
