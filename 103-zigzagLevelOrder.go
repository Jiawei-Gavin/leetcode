package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution1
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	zigzag(root, 0, &result)
	return result
}

func zigzag(root *TreeNode, depth int, result *[][]int) {
	if root == nil {
		return
	}
	for len(*result) < depth+1 {
		*result = append(*result, []int{})
	}
	if depth%2 == 0 {
		(*result)[depth] = append((*result)[depth], root.Val)
	} else {
		(*result)[depth] = append([]int{root.Val}, (*result)[depth]...)
	}
	zigzag(root.Left, depth+1, result)
	zigzag(root.Right, depth+1, result)
}

// solution2
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue, flag := []*TreeNode{root}, false
	for len(queue) > 0 {
		size := len(queue)
		var tmp []*TreeNode
		vals := make([]int, size)
		j := size - 1
		for i := 0; i < size; i++ {
			root = queue[0]
			queue = queue[1:]
			if !flag {
				vals[i] = root.Val
			} else {
				vals[j] = root.Val
				j--
			}
			if root.Left != nil {
				tmp = append(tmp, root.Left)
			}
			if root.Right != nil {
				tmp = append(tmp, root.Right)
			}
		}
		result = append(result, vals)
		flag = !flag
		queue = tmp
	}
	return result
}
