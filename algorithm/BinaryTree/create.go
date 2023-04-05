package binarytree

type ListNode struct {
	Data  int
	Left  *ListNode
	Right *ListNode
}

func NewListNode(data int) *ListNode {
	return &ListNode{Data: data}
}

func NewTree(root *ListNode, data int) {
	cur := root
	for cur != nil {
		if cur.Data > data {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = NewListNode(data)
				return
			}
		} else {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = NewListNode(data)
				return
			}
		}
	}
	fmt.Println(111)
}