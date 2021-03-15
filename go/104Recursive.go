/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    
    ret :=0
    if root == nil{
        return 0
    }else{
        ret = 1
    }
    
    if root.Left == nil && root.Right == nil{
        return 1
    }
    l,r := 0,0
    if root.Left != nil{
        l= maxDepth(root.Left)        
    }
    if root.Right != nil{
        r= maxDepth(root.Right)        
    }
    
    if l>r{
        return ret + l
    }else{
        return ret + r
    }
}
