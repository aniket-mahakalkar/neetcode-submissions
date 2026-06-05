/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {

	inorderMap := make(map[int]int)
	currRootInd := 0

	for i, val := range inorder {

		inorderMap[val] = i
	}

	var build func(left, right int) *TreeNode 

	build = func(left, right int) *TreeNode{

		if left > right {
			return nil
		}

		root := &TreeNode{}
		root.Val = preorder[currRootInd]

		currRootInd++

		mid := inorderMap[root.Val]

		root.Left = build(left, mid-1)
		root.Right = build(mid+1,right)

		return root
	}
    

	return build(0,len(inorder)-1)
}
