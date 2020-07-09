/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    ret := &ListNode{}
    var x int
    var y int
    retHead := ret
    
    if(l1 == nil){
        return l2
    }
    if(l2 == nil){
        return l1
    }
    
    t1 := l1
    t2 := l2
    
    for{
        if(t1 == nil && t2 == nil){
            ret = nil
            break
        }
        
        if(t1 == nil){
            ret.Val=t2.Val
            t2=t2.Next
            if(t2!=nil){
                ret.Next = &ListNode{}
                ret=ret.Next
                continue
            }else{
                ret.Next = nil
                break
            }
            
            
        }
        
        if(t2 == nil){
            ret.Val=t1.Val
            t1=t1.Next
            if(t1!=nil){
                ret.Next = &ListNode{}
                ret=ret.Next
                continue
            }else{
                ret.Next = nil
                break
            }
            
        }
        
        x = t1.Val
        y = t2.Val
        
        if(x<y){
            ret.Val=x
            t1 = t1.Next
        }else{
            ret.Val=y
            t2 = t2.Next
        }
        ret.Next = &ListNode{}
        ret = ret.Next
    }
    return retHead
    
}
