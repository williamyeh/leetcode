func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    t1 := l1
    t2 := l2
    temp := &ListNode{}
    res := temp
    var x, y, carry int
    carry = 0
    
    
    
    for{

        
        if(t1==nil && t2==nil ){
            if(carry == 1){
                temp.Val = 1
                temp.Next = nil    
            }
            break
        }
        
        if(t1 != nil){
            x = t1.Val
            t1 = t1.Next
        }else{
            x=0
        }
        
        if(t2 != nil){
            y = t2.Val
            t2 = t2.Next
        }else{
            y=0
        }
        
        sum := x+y+carry
        carry = sum/10
        remain := sum%10
        
        temp.Val = remain
        
        if(t1==nil && t2==nil && carry==0){
            temp.Next = nil
            break
        }
        temp.Next = &ListNode{}
        temp = temp.Next
    }
    return res
}
