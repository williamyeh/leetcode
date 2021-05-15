func myPow(x float64, n int) float64 {
    
    var ret float64
    if x == 1.0{
        return 1.0
    }
    if x == -1.0{
        if n%2 == 0{
            return 1.0
        }else{
            return -1.0
        }
        
    }
    if(n==0){
        return 1.0
    }else if(n>0){
        ret = x
    }else{
        ret = 1/x
        x = 1/x
        n = -n
    }
    // fmt.Println(ret,n)

    i:=1    
    for ;2*i<n ; i+=i {
        ret *= ret
    }
    
    for j:=n-i ; j>0 ; j--{
        ret *=x
    }
    return ret
    
}
