package ex01

import (
	"container/list"
	"day5/tree"
)

func UnrollGarland(root *tree.TreeNode) []bool {
	var layers []*list.List
	walk(&layers, 0, root)
	l := 0
	for k := len(layers) - 1; k != -1; k-- {
		l += (layers[k].Len())
	}
	ret := make([]bool, l)
	l = 0
	for _, i := range layers {
		for j := i.Front(); j != nil; j = j.Next() {
			ret[l] = j.Value.(bool)
			l++
		}
	}
	return ret
}

func walk(layers *[]*list.List, k int, root *tree.TreeNode) {
	if root == nil {
		return
	}
	if len(*layers) <= k {
		*layers = append(*layers, list.New().Init())
	}
	if k%2 == 0 {
		(*layers)[k].PushBack(root.HasToy)
	} else {
		(*layers)[k].PushFront(root.HasToy)
	}
	walk(layers, k+1, root.Right)
	walk(layers, k+1, root.Left)
}

// func UnrollGarland(root *tree.TreeNode) []bool {
// 	var layers [][]bool
// 	walk(&layers, 0, root)
// 	var ret []bool
// 	for k, i := range layers {
// 		if k % 2 == 0 {
// 			ret = append(ret, i...)
// 		} else {
// 			for j,_ := range i {
// 				ret = append(ret, i[len(i) - j - 1])
// 			}
// 		}
// 	}
// 	return ret
// }

// func walk(layers *[][]bool, k int, root *tree.TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	if len(*layers) <= k {
// 		*layers = append(*layers, []bool{})
// 	}
// 	(*layers)[k] = append((*layers)[k], root.HasToy)
// 	walk(layers, k + 1, root.Right)
// 	walk(layers, k + 1, root.Left)
// }
