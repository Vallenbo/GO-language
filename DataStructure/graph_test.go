package DataStructure

import (
	"fmt"
	"testing"
)

type Graph struct {
	Node1s []*Node1
}
type Node1 struct {
	Value int
	Edges []*Edge
}
type Edge struct {
	Weight int
	Node1  *Node1
}

func (g *Graph) AddNode1(n *Node1) {
	g.Node1s = append(g.Node1s, n)
}
func (n *Node1) AddEdge(e *Edge) {
	n.Edges = append(n.Edges, e)
}
func (g *Graph) Traverse() {
	for _, Node1 := range g.Node1s {
		fmt.Println("Node1:", Node1.Value)
		for _, edge := range Node1.Edges {
			fmt.Println("Edge:", edge.Weight, "to Node1:", edge.Node1.Value)
		}
	}
}

func Test_graph(T *testing.T) {
	graph := Graph{}
	Node11 := Node1{Value: 1}
	Node12 := Node1{Value: 2}
	Node13 := Node1{Value: 3}
	Node14 := Node1{Value: 4}
	edge1 := Edge{Weight: 1, Node1: &Node12}
	edge2 := Edge{Weight: 2, Node1: &Node13}
	edge3 := Edge{Weight: 3, Node1: &Node14}
	Node11.AddEdge(&edge1)
	Node11.AddEdge(&edge2)
	Node12.AddEdge(&edge3)
	graph.AddNode1(&Node11)
	graph.AddNode1(&Node12)
	graph.AddNode1(&Node13)
	graph.AddNode1(&Node14)
	graph.Traverse()
}

//注释：
//Graph 结构体表示整个图，包含一个节点切片 Node1s，表示所有的节点。
//Node1 结构体表示图中的一个节点，包含一个 int 类型的值 Value 和一个边切片 Edges，表示与该节点相连的所有边。
//Edge 结构体表示一条边，包含一个 int 类型的权重 Weight 和一个指向另一个节点的指针 Node1。
//AddNode1 方法用于向图中添加节点。
//AddEdge 方法用于向节点中添加边。
//Traverse 方法用于遍历整个图，输出每个节点以及与其相连的所有边。
//在 main 函数中，我们创建了一个图，四个节点，以及三条边，然后将它们加入到图中，最后遍历整个图并输出。
