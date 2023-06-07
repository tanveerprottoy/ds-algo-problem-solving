package binary

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, res)
	return res
}

func inorder(node *TreeNode, res []int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	res = append(res, node.Val)
	inorder(node.Right, res)
}
