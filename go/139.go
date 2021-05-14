func wordBreak(s string, wordDict []string) bool {
    
    lens := len(s)
    dp := make([]bool,lens+1,lens+1)
    hashmap := map[string]bool{}
    
    for i:=0 ; i<len(wordDict) ; i++{
        hashmap[wordDict[i]]=true
    }
    trues := []int{}
    
    for i,_ := range s{
        // fmt.Printf("%c %b \n",v,dp[i])
       
        // fmt.Println(s[0:i+1])
        if(hashmap[s[0:i+1]]){
            dp[i+1] = true
            trues = append(trues,i+1)
        }else{
            for j:=0; j<len(trues) ; j++{
                // fmt.Println(trues)
                if(trues[j]<=i+1 && hashmap[s[trues[j]:i+1]]){
                    dp[i+1] = true
                    trues = append(trues,i+1)
                    break
                }
            }
        }
    }
    // fmt.Println(dp)
    return dp[lens]
}
