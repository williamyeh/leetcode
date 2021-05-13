var count int

func convertBST(root *TreeNode) *TreeNode {
    
    if root == nil{
        return &TreeNode{}
    }
    ret := &TreeNode{Val:0 , Left:nil, Right:nil}
    rightmidleft(root,ret)
    return ret
}

func rightmidleft(root *TreeNode, ret *TreeNode){
    if root.Right != nil{
        ret.Right = &TreeNode{Val:0 , Left:nil, Right:nil}
        rightmidleft(root.Right,ret.Right)
    }
    // fmt.Println(root.Val)
    ret.Val = root.Val+count
    count =ret.Val
    if root.Left !=nil{
        ret.Left = &TreeNode{Val:0 , Left:nil, Right:nil}
        rightmidleft(root.Left,ret.Left)
    }
}
