package main

import "fmt"

//Graph Structure
type graph struct {
	vertices []*vertex
}

//Vertex Structure
type vertex struct {
	key      int
	adjacent []*vertex
}

//add vertex
func (g *graph) addVertex(k int) {
	g.vertices = append(g.vertices, &vertex{key: k})
}

//print
func (g *graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)

		for _, v := range g.vertices {
			fmt.Printf(" %v ", v.key)
		}
	}
}
func main() {
	test := &graph{}

	for i := 0; i < 5; i++ {
		test.addVertex(i)
	}
	test.Print()
}
