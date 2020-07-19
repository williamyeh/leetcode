func mySqrt(x int) int {
    
    if(x < 2){
        return x
    }
    
    low := 0
    high := x
    mid := 0
    square := 0
    
    for{
        mid = (low+high)/2
        if(low == mid){
            return mid
        }
        
        square = mid*mid
        if(square == x){
            return mid
        }
        
        if(square > x){
            high = mid
        }else{
            low = mid
        }
    }
    return low
    
}
