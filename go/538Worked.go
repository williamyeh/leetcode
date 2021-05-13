func convertBST(root *TreeNode) *TreeNode {
    count := 0
    if root == nil{
        return nil
    }
    rightmidleft(root, &count)
    return root
}

func rightmidleft(root *TreeNode, count *int){
    
    if root == nil{
        return
    }
    rightmidleft(root.Right,count)
    root.Val += *count
    *count = root.Val
    rightmidleft(root.Left,count)
    
}
