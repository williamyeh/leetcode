func nextPermutation(nums []int)  {
    
    var dec int
    var second int
    // var start int
    
    // find the dec point
    for i:=len(nums)-1 ; i>0 ; i--{
        if(nums[i]<=nums[i-1]){
            continue
        }else{
            dec = i-1
            break
        }
    }
    
    // find the second
    for i:=len(nums)-1 ; i> dec ; i--{
        if(nums[i]>nums[dec]){
            second = i
            break
        }
    }
    
    if(second == 0 && dec == 0){
        for i,j := 0,len(nums)-1 ; i<j ; i,j = i+1,j-1{
        // fmt.Println(nums,i,j)
            nums[i] , nums[j] = nums[j] , nums[i]
        // fmt.Println(nums,i,j)
        }
        return
    }
    // fmt.Println(dec,second)
    nums[dec],nums[second] =  nums[second],nums[dec]
    fmt.Println(nums)
    // swap nums[dec+1:len(nums)]
    for i,j := dec+1,len(nums)-1 ; i<j ; i,j = i+1,j-1{
        // fmt.Println(nums,i,j)
        nums[i] , nums[j] = nums[j] , nums[i]
        // fmt.Println(nums,i,j)
    }
    
}
