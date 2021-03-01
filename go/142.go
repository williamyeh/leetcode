func detectCycle(head *ListNode) *ListNode {

    slow:=head
    fast:=head
    flag:=false
    for fast !=nil && fast.Next!=nil{
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast{
            flag=true
            break
        }
    }
    if(flag){
        slow=head
        for fast != slow{
            slow = slow.Next
            fast = fast.Next
        }
        return slow
    }else{
        return nil;    
    }
}
