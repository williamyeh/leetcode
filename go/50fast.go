func myPow(x float64, n int) float64 {
    
    if(n==0){
        return 1
    }
    
    if n<0{
        n *= -1
        x= 1/x
    }
    ret,base := 1.0, x
    
    for n>1{
        if(n%2 == 0){
            n = n >> 1
            base *= base
            continue            
        }
        n -=1
        ret *= base
    }
    
    return ret*base
    
}
