func hasCycle(head *ListNode) bool {
    
    
    hmap := map[*ListNode]bool{}
    var temp *ListNode
    temp = new(ListNode)
    temp.Next = head
    
    if(head == nil || head.Next == nil){
        return false
    }else{
        for temp=temp.Next ; temp!=nil ; temp=temp.Next{
            if _, in := hmap[temp]; in{
                return true
            }else{
                hmap[temp] = true
            }
        }    
    }
    return false
    
    
}
