package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for l != nil {
		if count == pos {
			return l
		}
		count++
		l = l.Next
	}
	return nil
}
