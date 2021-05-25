func lengthOfLongestSubstring(s string) int {
    
    var max int
    var hashmap map[rune]bool
    // fmt.Println(len(s))
    for i:=0; i<len(s) ; i++{
        hashmap = make(map[rune]bool)
        counter :=0
        for _,v := range(s[i:len(s)]){
            // fmt.Println(v)
            if _,ok := hashmap[v]; !ok{
                hashmap[v] = true
                counter++
            }else{
                if(counter > max){
                    max = counter
                }
                break
            }
            if(counter > max){
                    max = counter
            }
        }
    }
    
    return max
}
