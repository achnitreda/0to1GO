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

// func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}
// 	BTreeApplyPostorder(root.Left, f)
// 	BTreeApplyPostorder(root.Right, f)
// 	_, err := f(root.Data)
// 	if err != nil {
// 		return
// 	}
// }
