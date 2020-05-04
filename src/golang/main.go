package main

import "fmt"

func main() {
	fmt.Println("Hello algorithm")
	//l1 := new(ListNode)
	//fmt.Println(l1.Val)
	tree := TreeNode{}
	tree.Val = 10
	tree.Left = &TreeNode{}
	tree.Right = &TreeNode{}
	tree.Left.Val = 9
	tree.Right.Val = 11
	tree.Left.Left = &TreeNode{}
	tree.Left.Left.Val = 8
	tree.Left.Left.Right = &TreeNode{}
	tree.Left.Left.Right.Val = 15
	//postorderTraversalStack(&tree)
	maxDepth(&tree)
}

//https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	
	dic := make(map[int]int)
	ret := make([]int, 2)
	
	for index, value := range nums {
		dicValue, ok := dic[value]
		if ok {
			ret[0] = dicValue
			ret[1] = index
			break
		} else {
			dic[target-value] = index
		}
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val1 := 0
	val2 := 0
	addtion := 0
	sum := 0
	newVal := 0
	ret := &ListNode{}
	tmp := ret
	
	for l1 != nil || l2 != nil {
		
		if l1 != nil {
			val1 = l1.Val
		} else {
			val1 = 0
		}
		
		if l2 != nil {
			val2 = l2.Val
		} else {
			val2 = 0
		}
		sum = val1 + val2 + addtion
		newVal = sum % 10
		addtion = sum / 10
		tmp.Val = newVal
		
		if (l1 != nil && l1.Next != nil) || (l2 != nil && l2.Next != nil) || addtion != 0 {
			tmp.Next = &ListNode{addtion, nil}
			tmp = tmp.Next
		}
		
		if l1 != nil {
			l1 = l1.Next
		}
		
		if l2 != nil {
			l2 = l2.Next
		}
		
	}
	return ret	
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int
var nums []int
var stack []*TreeNode

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	return perorderTraversalStack(root);
}

func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	res = inorderTraversalStack(root)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	res = []int{}
	postorderTraversalRecursive(root)
	return res
}

//迭代前序遍历
func perorderTraversalStack(root *TreeNode) []int {
	res = []int{}         //存储结果
	stack = []*TreeNode{} //存储每一个节点的右子节点的指针
	
	for len(stack) > 0 || root != nil {
		//将所有的左节点遍历并输出，并把包含的右子节点加入到stack中
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		//从stack尾部取出节点
		index := len(stack) - 1
		root = stack[index]
		stack = stack[:index]
	}
	return res
}

//迭代中序遍历
func inorderTraversalStack(root *TreeNode) []int {
	res = []int{}         //存储结果
	stack = []*TreeNode{} //存储每一个节点的右子节点的指针
	
	for len(stack) > 0 || root != nil {
		//将所有左子节点全部放入栈中
		for root != nil {
			
			stack = append(stack, root)
			root = root.Left
		}
		//从stack尾部取出节点，并输出第一个节点
		//该节点存在右子节点放到stack尾部
		index := len(stack) - 1
		stack = stack[:index]
		res = append(res, root.Val)
		if stack[index] != nil {
			root = root.Right
		}
	}
	return res
}

//迭代后序遍历
func postorderTraversalStack(root *TreeNode) []int {
	res = []int{}             //存储结果
	stack = []*TreeNode{root} //存储每一个节点的右子节点的指针
	//stack = append(stack, root)
	for len(stack) > 0 && root != nil {
		index := len(stack) - 1
		//stack = append(stack, root)
		if root.Left != nil || root.Right != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
				root.Right = nil
				index++
			}
			
			if root.Left != nil {
				stack = append(stack, root.Left)
				root.Left = nil
				index++
			}
			
		} else if root.Left == nil && root.Right == nil {
			res = append(res, root.Val)
			stack = stack[:index]
			index--
		}
		if index >= 0 {
			root = stack[index]
		}
	}
	return res
}
func postorderTraversalStack1(root *TreeNode) []int {
	res = []int{}             //存储结果
	stack = []*TreeNode{root} //存储每一个节点的右子节点的指针
	//stack = append(stack, root)
	for len(stack) > 0 && root != nil {
		//将所有左子节点全部放入栈中
		index := len(stack) - 1
		root = stack[index]
		
		if root.Right != nil {
			stack = append(stack, root.Right)
			index++
		}
		
		if root.Left != nil {
			stack = append(stack, root.Left)
			index++
		}
		
		if root.Left == nil && root.Right == nil {
			//index := len(stack) - 1
			res = append(res, root.Val)
			stack = stack[:index]
		}
		//入栈过的左右子节点置为空
		root.Left = nil
		root.Right = nil
	}
	
	return res
}

//递归前序遍历
func perorderTraversalRecursive(root *TreeNode) {
	
	if root != nil {
		nums = append(nums, root.Val)
		perorderTraversalRecursive(root.Left)
		perorderTraversalRecursive(root.Right)
		
	}
}

//递归中序遍历
func inorderTraversalRecursive(root *TreeNode) {
	
	if root != nil {
		perorderTraversalRecursive(root.Left)
		nums = append(nums, root.Val)
		perorderTraversalRecursive(root.Right)
	}
}

//递归后序遍历
//https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func postorderTraversalRecursive(root *TreeNode) {
	
	if root != nil {
		perorderTraversalRecursive(root.Left)
		perorderTraversalRecursive(root.Right)
		nums = append(nums, root.Val)
	}
}

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	//max := findmaxDepth(root)
	return findmaxDepth(root)
}

func findmaxDepth(root *TreeNode) int {
	
	if root == nil {
		return 0
	}else{
		//找左子树
		left := findmaxDepth(root.Left)
		//找右子树
		right := findmaxDepth(root.Right)
		
		if left > right {
			return left+1
		}else{
			return right + 1
		}
	}
	
	
}
