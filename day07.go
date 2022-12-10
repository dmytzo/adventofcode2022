package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type node struct {
	dir      bool
	name     string
	children map[string]*node
	parent   *node
	size     int
}

func (n *node) nodeSize() int {
	if !n.dir {
		return n.size
	}

	var size int

	for _, v := range n.children {
		size += v.nodeSize()
	}

	return size
}

func getTree(input []string) *node {
	var (
		head        *node
		currentNode *node
	)

	for _, line := range input {
		tokens := strings.Split(line, " ")

		switch tokens[0] {
		case "$":
			if tokens[1] != "cd" {
				continue
			}

			if tokens[2] == ".." {
				currentNode = currentNode.parent
				continue
			}

			n := &node{
				dir:      true,
				name:     tokens[2],
				children: make(map[string]*node),
			}

			if currentNode == nil {
				currentNode = n
				head = n
				continue
			}

			n.parent = currentNode
			currentNode.children[n.name] = n

			currentNode = n
		case "dir":
			currentNode.children[tokens[1]] = &node{
				dir:      true,
				name:     tokens[1],
				parent:   currentNode,
				children: make(map[string]*node),
			}
		default:
			size, err := strconv.Atoi(tokens[0])
			if err != nil {
				log.Fatalf(err.Error())
			}

			currentNode.children[tokens[1]] = &node{
				name:   tokens[1],
				parent: currentNode,
				size:   size,
			}
		}
	}

	return head
}

func getSum(d *node) int {
	var res int
	for _, v := range d.children {
		if !v.dir {
			continue
		}

		res += getSum(v)
		if v.nodeSize() <= 100000 {
			res += v.nodeSize()
		}
	}

	return res
}

func getDirSizes(d *node) []int {
	var res []int

	for _, v := range d.children {
		if v.dir {
			res = append(res, v.nodeSize())
			res = append(res, getDirSizes(v)...)
		}
	}

	return res
}

func daySevenTaskOne(input []string) int {
	return getSum(getTree(input))
}

func daySevenTaskTwo(input []string) int {
	var (
		head         = getTree(input)
		diskSpace    = 70000000
		targetSpace  = 30000000
		currentSpace = diskSpace - head.nodeSize()
		needToDelete = targetSpace - currentSpace
	)

	if needToDelete <= 0 {
		return 0
	}

	var res int
	for _, v := range getDirSizes(head) {
		if v >= needToDelete {
			if res == 0 || v < res {
				res = v
			}
		}
	}

	return res
}

func DaySeven() {
	input := inputLines("input/7.txt")
	fmt.Printf("### DAY 7 ###\n 1: %d\n 2: %d\n", daySevenTaskOne(input), daySevenTaskTwo(input))
}
