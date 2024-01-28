package main

//646. Maximum Length of Pair Chain
func findLongestChain(pairs [][]int) int {
	type node struct {
		me *[]int
		parent *[]int
		left *[]int
		right *[]int
	}
	var n node
	n.me = &pairs[0]
	n.parent = nil
	nodes := []node{}
	for _, v := range pairs {
		nodes = append(nodes, node{&v, nil, nil, nil})
	}
	// for _, input := range pairs {
	// 	for _, n := range nodes {
	// 		if input[0] == *n.left[] + 1 {

	// 		}
	// 	}
	// }
	return 0
}
