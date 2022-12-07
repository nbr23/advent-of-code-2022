package main

import (
	// "github.com/nbr23/advent-of-code-2022/utils/inputs"

	"strings"

	"github.com/nbr23/advent-of-code-2022/utils/inputs"
	"github.com/nbr23/advent-of-code-2022/utils/utils"
	//	"github.com/pkg/profile"
)

var day_num int = 7

type Node struct {
	parent   *Node
	children map[string]*Node // name -> Node
	size     int
	name     string
	isdir    bool
}

func propagateSizeUp(node *Node, size int) {
	if node == nil {
		return
	}
	node.size += size
	propagateSizeUp(node.parent, size)
}

func sumDirsUnder(dirList []*Node, maxsize int) int {
	res := 0
	for _, dir := range dirList {
		if dir.size <= maxsize {
			res += dir.size
		}
	}
	return res
}

func buildTree(input []string) (*Node, []*Node) {
	currentNode := &Node{
		nil,
		make(map[string]*Node),
		0,
		"/",
		true,
	}
	rootNode := currentNode
	previousNode := currentNode
	var dirList []*Node

	for i := 0; i < len(input); i++ {
		if input[i] == "$ ls" {
			for i = i + 1; i < len(input) && input[i][0] != '$'; i++ {
				lsline := strings.Split(input[i], " ")
				if lsline[0] == "dir" {
					dirnode := Node{
						currentNode,
						make(map[string]*Node),
						0,
						lsline[1],
						true,
					}
					currentNode.children[lsline[1]] = &dirnode
					dirList = append(dirList, &dirnode)
				} else {
					filenode := Node{
						currentNode,
						nil,
						inputs.ParseDecInt(lsline[0]),
						lsline[1],
						false,
					}
					currentNode.children[lsline[1]] = &filenode
					propagateSizeUp(currentNode, filenode.size)
				}
			}
			i--
		} else if input[i][0:4] == "$ cd" {
			if input[i] == "$ cd /" {
				previousNode = currentNode
				currentNode = rootNode
			} else if input[i] == "$ cd .." {
				previousNode, currentNode = currentNode, currentNode.parent
			} else {
				previousNode = currentNode
				currentNode = previousNode.children[strings.Split(input[i], " ")[2]]
			}
		} else {
			panic("shouldn't happen really")
		}
	}
	return rootNode, dirList
}

var root *Node
var dirList []*Node

func part1(input string) interface{} {
	root, dirList = buildTree(inputs.InputToStrList(input))
	return sumDirsUnder(dirList, 100000)
}

func part2(input string) interface{} {
	totalFS := 70000000
	neededFreeFS := 30000000
	freespace := totalFS - root.size
	need := neededFreeFS - freespace

	res := neededFreeFS

	for _, dir := range dirList {
		if dir.size >= need && dir.size <= res {
			res = dir.size
		}
	}
	return res
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, day_num)
}
