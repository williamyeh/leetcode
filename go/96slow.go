func numTrees(n int) int {
    var sum int
    switch(n){
    case 0: 
        return 1
        break
    case 1:
        return 1 
        break
    case 2:
        return 2
        break
    case 3:
        return 5 
        break
    default:
        
        for i:=1 ; i<=n ;i++{
            sum += numTrees(i-1)*numTrees(n-i)
        }
     
    }
    
    return sum
}
