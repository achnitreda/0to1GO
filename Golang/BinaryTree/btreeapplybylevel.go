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

// func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
// 	height := BTreeLevelCount(root)

// 	for level := 1; level <= height; level++ {
// 		printLevel(root, level, f)
// 	}
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

// func printLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}
// 	if level == 1 {
// 		_, err := f(root.Data)
// 		if err != nil {
// 			return
// 		}
// 	} else if level > 1 {
// 		printLevel(root.Left, level-1, f)
// 		printLevel(root.Right, level-1, f)
// 	}
// }
