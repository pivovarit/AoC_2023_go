package day25

import (
	"github.com/pivovarit/aoc/util"
	"strings"
)

func run() {
	input := util.ReadInput()

	util.Timed("snowverloadPart1", func() int {
		return snowverloadPart1(input)
	})
}

func snowverloadPart1(input []string) int {
	groups := MultipleOfTwoGroups(input)
	firstGroup, secondGroup := groups.split(3)
	return len(firstGroup) * len(secondGroup)
}

type Node struct {
	name     string
	edges    map[*Edge]*Node
	traveled bool
}

func NewNode(name string) *Node { return &Node{name: name, edges: make(map[*Edge]*Node)} }

type Edge struct {
	visited bool
}

type Graph struct {
	nodes map[string]*Node
	edges []*Edge
}

func (g *Graph) removeShortestPath(source, dest *Node) bool {
	type QueueItem struct {
		edge     *Edge
		node     *Node
		previous *QueueItem
	}

	queue := make([]*QueueItem, 0, len(g.nodes))
	queue = append(queue, &QueueItem{node: source})

	found := false
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node == dest {
			for itr := current; itr.edge != nil; itr = itr.previous {
				itr.edge.visited = true
			}
			found = true
			break
		}

		for e, n := range current.node.edges {
			if e.visited || n.traveled {
				continue
			}
			n.traveled = true
			queue = append(queue, &QueueItem{e, n, current})
		}
	}
	for _, n := range g.nodes {
		n.traveled = false
	}
	return found
}

func (g *Graph) cutPaths(source, dest *Node, pathNum int) bool {
	complete := true
	for i := 0; i < pathNum; i++ {
		if !g.removeShortestPath(source, dest) {
			complete = false
			break
		}
	}
	for _, e := range g.edges {
		e.visited = false
	}
	return complete
}

func (g *Graph) split(cuts int) ([]*Node, []*Node) {
	g1, g2 := []*Node{}, []*Node{}

	var source *Node
	for _, n := range g.nodes {
		source = n
		break
	}
	g1 = append(g1, source)

	for _, dest := range g.nodes {
		if source == dest {
			continue
		}

		if g.cutPaths(source, dest, cuts+1) {
			g1 = append(g1, dest)
		} else {
			g2 = append(g2, dest)
		}
	}
	return g1, g2
}

func MultipleOfTwoGroups(input []string) *Graph {
	nodes := make(map[string]*Node)
	for _, line := range input {
		name, _, _ := strings.Cut(line, ": ")
		nodes[name] = NewNode(name)
	}
	edges := make([]*Edge, 0)
	for _, line := range input {
		sourceName, destNames, _ := strings.Cut(line, ": ")
		source := nodes[sourceName]
		for _, destName := range strings.Split(destNames, " ") {
			dest, ok := nodes[destName]
			if !ok {
				dest = NewNode(destName)
				nodes[destName] = dest
			}
			newEdge := &Edge{}
			edges = append(edges, newEdge)
			source.edges[newEdge] = dest
			dest.edges[newEdge] = source
		}
	}
	return &Graph{nodes, edges}
}
