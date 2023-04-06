package tree

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
)

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewListNode(data int) *TreeNode {
	return &TreeNode{Val: data}
}
func (l *TreeNode) NewTree(n *TreeNode, data int) bool {
	cur := n
	for cur != nil {
		if cur.Val > data {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = NewListNode(data)
				return true
			}
		} else {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = NewListNode(data)
				return true
			}
		}
	}

	return false
}

// PreorderTraversal 前序
func PreorderTraversal(root *TreeNode) []int {
	var res []int
	preorder := func(r *TreeNode) {}
	preorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		res = append(res, r.Val)
		preorder(r.Left)
		preorder(r.Right)
	}
	preorder(root)
	return res
}

// PreorderTraversal2 前序非递归
func PreorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// InOrder 中序
func InOrder(root *TreeNode) []int {
	var res []int
	io := func(r *TreeNode) {}
	io = func(r *TreeNode) {
		if r == nil {
			return
		}
		io(r.Left)
		res = append(res, r.Val)
		io(r.Right)
	}
	io(root)
	return res
}

// InOrder2 中序 非递归
func InOrder2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

// PostOrder 后续遍历
func PostOrder(root *TreeNode) []int {
	var res []int
	po := func(r *TreeNode) {}
	po = func(r *TreeNode) {
		if r == nil {
			return
		}
		po(r.Left)
		po(r.Right)
		res = append(res, r.Val)
	}
	po(root)
	return res
}

// PostOrder2 后续 非递归
func PostOrder2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var node *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right != nil && root.Right != node {
			stack = append(stack, root)
			root = root.Right
		} else {
			res = append(res, root.Val)
			node = root
			root = nil
		}
	}
	return res
}

// LevelOrder 层序遍历 递归
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	depth := 0
	order := func(r *TreeNode, depth int) {}
	order = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], r.Val)

		order(r.Left, depth+1)
		order(r.Right, depth+1)
	}

	order(root, depth)
	return res
}

// LevelOrder2 层序遍历 非递归 队列
func LevelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var tmp []int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		num := queue.Len()
		for i := 0; i < num; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}

// LevelOrder3 层序遍历 非递归 数组
func LevelOrder3(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var curStack []*TreeNode
	curStack = append(curStack, root)
	for len(curStack) > 0 {
		var tmp []int
		var nextStack []*TreeNode
		for _, node := range curStack {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				nextStack = append(nextStack, node.Left)
			}
			if node.Right != nil {
				nextStack = append(nextStack, node.Right)
			}
		}
		res = append(res, tmp)
		curStack = nextStack
	}

	return res
}

// 层序遍历 自底向上 层序之后反转

// RightSideView 二叉树的右视图  取每层的最后一个元素 递归深度
func RightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	dfs := func(r *TreeNode, depth int) {}
	dfs = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}
		if len(res) == depth {
			res = append(res, r.Val)
		}
		dfs(r.Right, depth+1)
		dfs(r.Left, depth+1)
	}
	dfs(root, 0)
	return res
}

// RightSideView2 二叉树的右视图  取每层的最后一个元素 非递归广度
func RightSideView2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var curStack []*TreeNode
	curStack = append(curStack, root)
	for len(curStack) > 0 {
		var nextStack []*TreeNode
		for _, node := range curStack {
			if node.Left != nil {
				nextStack = append(nextStack, node.Left)
			}
			if node.Right != nil {
				nextStack = append(nextStack, node.Right)
			}
		}
		res = append(res, curStack[len(curStack)-1].Val)
		curStack = nextStack
	}

	return res
}

// InvertTree 翻转二叉树 前序 递归
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

// InvertTree2 翻转二叉树 前序 非递归
func InvertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var stack []*TreeNode
	node := root
	stack = append(stack, node)
	for len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return root
}

// InvertTree3 翻转二叉树 前序 层序
func InvertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var stack []*TreeNode
	node := root
	stack = append(stack, node)
	for len(stack) > 0 {
		var nextStack []*TreeNode
		for i := 0; i < len(stack); i++ {
			stack[i].Left, stack[i].Right = stack[i].Right, stack[i].Left
			if stack[i].Left != nil {
				nextStack = append(nextStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				nextStack = append(nextStack, stack[i].Right)
			}
		}
		stack = nextStack
	}

	return root
}

// IsSymmetric 对称二叉树 递归
func IsSymmetric(root *TreeNode) bool {
	var dfs func(l,r *TreeNode) bool
	dfs = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		return dfs(l.Left, r.Right) && dfs(r.Left, l.Right)
	}

	return dfs(root.Left, root.Right)
}

// IsSymmetric2 对称二叉树 非递归
func IsSymmetric2(root *TreeNode) bool {
	var stack []*TreeNode
	stack = append(stack, root.Left, root.Right)

	for len(stack) > 0 {
		l := stack[0]
		r := stack[1]
		stack = stack[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil || l.Val != r.Val {
			return false
		}
		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
	}

	return true
}

// MaxDepth 最大深度 递归
func MaxDepth(root *TreeNode) int {
	var num int
	dfs := func(r *TreeNode, depth int) {}
	dfs = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}
		if depth > num {
			num = depth
		}
		dfs(r.Left, depth+1)
		dfs(r.Right, depth+1)
	}

	dfs(root, 1)
	return num
}

// MaxDepth2 最大深度 层序 栈
func MaxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var num int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		fmt.Println(stack)
		var nextStack []*TreeNode
		for i := 0; i < len(stack); i++ {
			if stack[i].Left != nil {
				nextStack = append(nextStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				nextStack = append(nextStack, stack[i].Right)
			}
		}
		if len(nextStack) > 0 {
			num++
		}
		stack = nextStack
	}
	return num+1
}


// MinDepth 最小深度 递归
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right == nil {
		return MinDepth(root.Left)+1
	}

	if root.Left == nil && root.Right != nil {
		return MinDepth(root.Right)+1
	}

	return minNum(MinDepth(root.Left), MinDepth(root.Right))+1
}

// MinDepth2 最小深度 层序
func MinDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	num := 1
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		num++
		l := len(stack)
		for i := 0; i < l; i++ {
			if stack[0].Left == nil && stack[0].Right == nil {
				return num
			}
			if stack[0].Left != nil {
				stack = append(stack, stack[0].Left)
			}
			if stack[0].Right != nil {
				stack = append(stack, stack[0].Right)
			}
			stack = stack[1:]
		}
	}

	return num
}

func minNum(a,b int) int {
	if a > b {
		return b
	}
	return a
}

// CountNodes 完全二叉树的节点个数
func CountNodes(root *TreeNode) int {
	var res int
	dfs := func(r *TreeNode) {}
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		res++
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return res
}

// CountNodes2 完全二叉树的节点个数
func CountNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		l := len(stack)
		res += l
		for i := 0; i < l; i++ {
			if stack[0].Left != nil {
				stack = append(stack, stack[0].Left)
			}
			if stack[0].Right != nil {
				stack = append(stack, stack[0].Right)
			}
			stack = stack[1:]
		}
	}

	return res
}

// IsBalanced 平衡二叉树 递归
func IsBalanced(root *TreeNode) bool {
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		lh := dfs(r.Left)
		rh := dfs(r.Right)
		if lh == -1 || rh == -1 {
			return -1
		}
		if lh - rh > 1 || rh - lh > 1 {
			return -1
		}

		return max(lh, rh) +1
	}

	if dfs(root) == -1 {
		return false
	}

	return true
}

// IsBalanced2 平衡二叉树 非递归
func IsBalanced2(root *TreeNode) bool {


	return true
}

func max(lh int, rh int) int {
	if lh > rh {
		return lh
	}

	return rh
}

// BinaryTreePaths 二叉树的所有路径 递归
func BinaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}

	dfs := func(r *TreeNode, tmp string) {}
	dfs = func(r *TreeNode, tmp string) {
		if r == nil {
			return
		}

		if tmp == "" {
			tmp = strconv.Itoa(r.Val)
		} else {
			tmp = tmp + "->" + strconv.Itoa(r.Val)
		}
		if r.Left == nil && r.Right == nil {
			res = append(res, tmp)
		}
		dfs(r.Left, tmp)
		dfs(r.Right, tmp)
	}

	dfs(root, "")

	return res
}

// BinaryTreePaths2 二叉树的所有路径 非递归
func BinaryTreePaths2(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}

	var stack []*TreeNode
	var paths []string
	stack = append(stack, root)
	paths = append(paths, "")
	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]
		path := paths[l-1]
		paths = paths[:l-1]

		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}

	return res
}

// SumOfLeftLeaves 左叶子之和 递归
func SumOfLeftLeaves(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	dfs := func(r *TreeNode) {}
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		if r.Left != nil {
			dfs(r.Left)
			if r.Left.Left == nil && r.Left.Right == nil {
				res += r.Left.Val
			}
		}
		if r.Right != nil {
			dfs(r.Right)
		}
	}

	dfs(root)
	return res
}

// SumOfLeftLeaves2 左叶子之和 非递归
func SumOfLeftLeaves2(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
			if node.Left.Left == nil && node.Left.Right == nil {
				res += node.Left.Val
			}
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return res
}

// FindBottomLeftValue 找树左下角的值
func FindBottomLeftValue(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	var tmpDepth int
	dfs := func(r *TreeNode, depth int) {}
	dfs = func(r *TreeNode, depth int) {
		if r== nil {
			return
		}

		dfs(r.Left, depth+1)

		dfs(r.Right, depth+1)
		if r.Left == nil && r.Right == nil {
			if tmpDepth < depth {
				tmpDepth = depth
				res = r.Val
			}
		}
	}

	dfs(root, 1)

	return res
}

// FindBottomLeftValue2 找树左下角的值 深度
func FindBottomLeftValue2(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		l := len(stack)

		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return res
}

// HasPathSum 路径总和 递归
func HasPathSum(root *TreeNode, targetSum int) bool {
	res := false
	if root == nil {
		return res
	}

	dfs := func(r *TreeNode, sum int) {}
	dfs = func(r *TreeNode, sum int) {
		if r == nil {
			return
		}
		sum += r.Val
		fmt.Println(sum)
		if sum == targetSum && r.Left == nil && r.Right == nil {
			res = true
			return
		}

		if r.Left != nil {
			dfs(r.Left, sum)

		}
		if r.Right != nil {
			dfs(r.Right, sum)
		}
	}


	dfs(root, 0)

	return res
}

// HasPathSum2 路径总和 非递归
func HasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var stack []*TreeNode
	stack = append(stack, root)
	sums := []int{root.Val}

	for len(stack) > 0 {
		l := len(stack)

		for i := 0; i < l; i++ {
			fmt.Println(sums)
			node := stack[0]
			sum := sums[0]
			stack = stack[1:]
			sums = sums[1:]
			if node.Left == nil && node.Right == nil && sum == targetSum {
				return true
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
				sums = append(sums, sum+node.Left.Val)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
				sums = append(sums, sum+node.Right.Val)
			}
		}

	}


	return false
}

// PathSum 路径总和 II
func PathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	dfs := func(r *TreeNode, nodes []int, ts int) {}
	dfs = func(r *TreeNode, nodes []int, ts int) {
		if r == nil {
			return
		}

		nodes = append(nodes, r.Val)
		ts -= r.Val
		if r.Left == nil && r.Right == nil && ts == 0 {
			res = append(res, append([]int{}, nodes...))
		}
		dfs(r.Left, nodes, ts)
		dfs(r.Right, nodes, ts)
	}

	dfs(root, []int{}, targetSum)

	return res
}

// BuildTree 从中序与后序遍历序列构造二叉树
func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}

	if len(postorder) == 1 {
		return root
	}
	index := 0
	for index = 0; index < len(inorder); index++ {
		if inorder[index] == root.Val {
			break
		}
	}
	fmt.Println(inorder, postorder, index)
	root.Left = BuildTree(inorder[:index], postorder[:index])
	root.Right = BuildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}

// BuildTree2 从中序与前序遍历序列构造二叉树
func BuildTree2(inorder []int, invertOrder []int) *TreeNode {
	if len(invertOrder) == 0 {
		return nil
	}

	root := &TreeNode{Val: invertOrder[0]}

	if len(invertOrder) == 1 {
		return root
	}
	index := 0
	for index = 0; index < len(inorder); index++ {
		if inorder[index] == root.Val {
			break
		}
	}
	fmt.Println(inorder, invertOrder, index)
	root.Left = BuildTree2(inorder[:index], invertOrder[1:index+1])
	root.Right = BuildTree2(inorder[index+1:], invertOrder[index+1:])

	return root
}

// ConstructMaximumBinaryTree 最大二叉树
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	num := 0
	index := 0
	for i, v := range nums {
		if v > num {
			num = v
			index = i
		}
	}

	root := &TreeNode{Val: num}

	root.Left = ConstructMaximumBinaryTree(nums[:index])
	root.Right = ConstructMaximumBinaryTree(nums[index+1:])

	return root
}

// MergeTrees 合并二叉树
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val

	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)

	return root1
}

// SearchBST 二叉搜索树中的搜索
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > val {
		return SearchBST(root.Left, val)
	} else {
		return SearchBST(root.Right, val)
	}
}

// SearchBST2 二叉搜索树中的搜索
func SearchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	for root != nil {
		if root.Val == val {
			return root
		}

		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}

// IsValidBST 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root, math.MinInt64, math.MaxInt64)
}

func check(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}

	if min >= int64(root.Val) || max <= int64(root.Val) {
		return false
	}

	// 对于左子树来说 当前节点是最大值，对于右子树来说当前节点是最小值
	return check(root.Left, min, int64(root.Val)) && check(root.Right, int64(root.Val), max)
}

// IsValidBST2 验证二叉搜索树 中序遍历
func IsValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 使用中序（从左边开始），从最底层开始遍历，可以保证下层的节点一定比上一层的节点小
	var pre *TreeNode
	var dfs func(r *TreeNode) bool
	dfs = func(r *TreeNode) bool {
		if r == nil {
			return true
		}

		lres := dfs(r.Left)
		if pre != nil && pre.Val >= r.Val {
			return false
		}
		pre = r
		rres := dfs(r.Right)
		return lres && rres
	}

	return dfs(root)
}

// GetMinimumDifference 二叉搜索树的最小绝对差
func GetMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var pre *TreeNode
	res := math.MaxInt64
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		dfs(r.Left)
		if pre != nil && r.Val - pre.Val < res {
			res = r.Val - pre.Val
		}
		pre = r
		dfs(r.Right)
	}

	dfs(root)
	return res
}

// FindMode 二叉搜索树中的众数
func FindMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var pre *TreeNode
	var dfs func(r *TreeNode)
	count := 1
	max := 1
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}

		dfs(r.Left)
		if pre != nil && pre.Val == r.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{r.Val}
			} else {
				res = append(res, r.Val)
			}
			max = count
		}
		pre = r
		dfs(r.Right)
	}
	dfs(root)

	return res
}

// LowestCommonAncestor 二叉树的最近公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	}

	return nil
}

// LowestCommonAncestor2 二叉搜索树的最近公共祖先
func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return LowestCommonAncestor2(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return LowestCommonAncestor2(root.Right, q, p)
	} else {
		return root
	}
}

// DeleteNode 删除二叉搜索树中的节点
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	
	if root.Val == key {
		// 左右为空
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 右为空
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		// 左为空
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		// 左右不空
		if root.Left != nil && root.Right != nil {
			// 整个左节点放到右节点的最左侧的节点上
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}

	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	}

	return root
}

// TrimBST 修剪二叉搜索树
func TrimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low || root.Val > high {
		// 左右为空
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 右为空
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		// 左为空
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		// 左右不空
		if root.Left != nil && root.Right != nil {
			// 整个左节点放到右节点的最左侧的节点上
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	if root.Val < high {
		root.Left = TrimBST(root.Left, low, high)
	}
	if root.Val > low {
		root.Right = TrimBST(root.Right, low, high)
	}

	return root
}

// SortedArrayToBST 将有序数组转换为二叉搜索树
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var makeTree func([]int) *TreeNode
	makeTree = func(ints []int) *TreeNode {
		if len(ints) == 0 {
			return nil
		}

		mid := (len(ints))/2
		root := &TreeNode{Val: ints[mid]}

		root.Left = makeTree(ints[:mid])
		root.Right = makeTree(ints[mid+1:])

		return root
	}

	return makeTree(nums)
}

// SortedArrayToBST2 将有序数组转换为二叉搜索树 非递归
//func SortedArrayToBST2(nums []int) *TreeNode {
//
//}