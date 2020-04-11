package graph


type Node struct {
	value int
}
type Graph struct {
	Nodes []*Node // 顶点数量
	Edges map[Node][]*Node // 邻接矩阵
}
// 获得顶点的数量
func (g *Graph) GetNodeNum() int {
	return len(g.Nodes)
}
// 获得边的数量
func (g *Graph) GetEdgeNum() int {
	return len(g.Edges)
}
// 添加节点
func (g *Graph) AddNode(value int) {
	g.Nodes = append(g.Nodes, &Node{value:value})
}
// 添加边
func (g *Graph) AddEdge(m, n int) {
	node1 := Node{m}
	node2 := Node{n}
	g.Edges[node1] = append(g.Edges[node1], &node2)
	g.Edges[node2] = append(g.Edges[node2], &node1)
}
