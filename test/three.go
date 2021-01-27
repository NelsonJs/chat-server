package main

import "fmt"

func main() {
	node := Node{}
	if node.child == nil {
		fmt.Println("is nil")
		return
	} else {
		fmt.Println("no nil")
	}
	node.child["a"] = Parent()
	node.child["a"].name = "af"
	fmt.Println(node.child["a"].name)
}

type Node struct {
	name string
	end bool
	position int
	child map[string]*Node
}

func Parent() *Node {
	return &Node{
		child: make(map[string]*Node),
	}
}

func Insert(word string) {
	if word == ""{
		return
	}
	node := Parent()
	for _,v := range word {
		if _,ok := node.child[string(v)]; !ok {

		}
	}
}

