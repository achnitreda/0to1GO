package piscine

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	return root
// }

// func BTreeLevelCount(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	leftLevel := BTreeLevelCount(root.Left)
// 	rightLevel := BTreeLevelCount(root.Right)
// 	if leftLevel > rightLevel {
// 		return leftLevel + 1
// 	} else {
// 		return rightLevel + 1
// 	}
// }
