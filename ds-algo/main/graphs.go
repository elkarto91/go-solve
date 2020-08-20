package main

import (
	"fmt"
	"strconv"
)

type Graph struct {
	Edges []*Edge
	Nodes []*Node
}

type Edge struct {
	Parent *Node
	Child  *Node
	Cost   int
}

type Node struct {
	Name string
}

const Infinity = int(^uint(0) >> 1)

func (g *Graph) AddNode(node *Node) {
	var isPresent bool
	for _, n := range g.Nodes {
		if n == node {
			isPresent = true
		}
	}

	if !isPresent {
		g.Nodes = append(g.Nodes, node)
	}
}

func (g *Graph) AddEdge(parent, child *Node, cost int) {

	edge := &Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}
	g.Edges = append(g.Edges, edge)
	g.AddNode(parent)
	g.AddNode(child)
}

func main() {

	fmt.Println("Nodes")
	a := &Node{Name: "a"}
	b := &Node{Name: "b"}
	c := &Node{Name: "c"}
	d := &Node{Name: "d"}
	e := &Node{Name: "e"}
	f := &Node{Name: "f"}
	g := &Node{Name: "g"}

	fmt.Println("Edges")
	graph := Graph{}
	graph.AddEdge(a, c, 2)
	graph.AddEdge(a, b, 5)
	graph.AddEdge(c, b, 1)
	graph.AddEdge(c, d, 9)
	graph.AddEdge(b, d, 4)
	graph.AddEdge(d, e, 2)
	graph.AddEdge(d, g, 30)
	graph.AddEdge(d, f, 10)
	graph.AddEdge(f, g, 1)

	fmt.Println(" Graph")
	/*	for k,v := range graph.Edges{
		if graph.Nodes[k]!=nil{
			fmt.Println("Node ",graph.Nodes[k])
		}
		fmt.Println(" Edge List",v)
	}*/

	s := graph.String()
	fmt.Println(" Graph now ", s)
}

// String returns a string representation of the Graph
func (g *Graph) String() string {
	var s string

	s += "Nodes: "
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
	}

	s += "\n"

	s += "Edges:\n"
	for _, edge := range g.Edges {
		s += edge.Parent.Name + " -> " + edge.Child.Name + " = " + strconv.Itoa(edge.Cost)
		s += "\n"
	}
	s += "\n"

	return s
}
