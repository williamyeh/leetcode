/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    list := []int{}
    getlist(root,&list)
    return maketree(&list,0,len(list)-1)
}

func maketree(list *[]int, l int, r int) *TreeNode{
    
    if l>r{return nil}
    if l==r{
        return &TreeNode{Val:(*list)[l], Left:nil, Right:nil}
    }
    mid := (l+r)/2
    ret := TreeNode{Val:(*list)[mid], Left: maketree(list,l,mid-1), Right: maketree(list,mid+1,r)}
    
    return &ret
}

func getlist(root *TreeNode, list *[]int) {
    
    if root == nil{
        return
    }
    getlist(root.Left, list)
    myappend(list,root.Val)
    // fmt.Println(list)
    getlist(root.Right, list)
}

func myappend(list *[]int, val int){
    *list=append(*list,val)    
}
