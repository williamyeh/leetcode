func lengthOfLongestSubstring(s string) int {
    
    var left int
    var right int
    var max int
    // var counter int
    hashmap := map[byte]bool{}
    
    for right < len(s){
       
        if _, ok := hashmap[s[right]]; !ok{
            hashmap[s[right]]=true
            right++
        }else{
            delete(hashmap,s[left])
            left++
        }
        if right-left > max{
            // fmt.Println(right,left)
            max = right-left
        }
        
    }
    
    return max
}
