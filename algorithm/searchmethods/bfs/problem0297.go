package bfs

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (codec *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for i := 0; i < len(queue); i++ {
		if queue[i] != nil {
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
	}
	for len(queue) > 0 && queue[len(queue)-1] == nil {
		queue = queue[0 : len(queue)-1]
	}
	result := make([]string, 0, len(queue))
	for _, node := range queue {
		if node == nil {
			result = append(result, "#")
		} else {
			result = append(result, strconv.Itoa(node.Val))
		}
	}
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	if strings.Compare(data, "") == 0 {
		return nil
	}
	nodeValues := strings.Split(data, ",")
	val, _ := strconv.Atoi(nodeValues[0])
	root := &TreeNode{Val: val}
	queue := make([]*TreeNode, 1, len(nodeValues))
	queue[0] = root
	currentRootIndex := 0
	isLeftChild := true
	for i := 1; i < len(nodeValues); i++ {
		if strings.Compare(nodeValues[i], "#") != 0 {
			val, _ = strconv.Atoi(nodeValues[i])
			child := &TreeNode{Val: val}
			if isLeftChild {
				queue[currentRootIndex].Left = child
			} else {
				queue[currentRootIndex].Right = child
			}
			queue = append(queue, child)
		}
		if !isLeftChild {
			currentRootIndex++
		}
		isLeftChild = !isLeftChild
	}
	return root
}
