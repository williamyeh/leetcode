var count int

func convertBST(root *TreeNode) *TreeNode {
    
    if root == nil{
        return &TreeNode{}
    }
    rightmidleft(root)
    return root
}

func rightmidleft(root *TreeNode){
    if root.Right != nil{
        rightmidleft(root.Right)
    }
    // fmt.Println(root.Val)
    root.Val += count
    count = root.Val
    if root.Left !=nil{
        rightmidleft(root.Left)
    }
}
