func sprint(nums []int)string{
    return fmt.Sprintf("%q",nums)
}

func threeSum(nums []int) [][]int {
    
    sort.Ints(nums)
    var ret [][]int
    mymap := make(map[string]bool)
    for i:=0 ; i<len(nums) ; i++{
        for j,k := i+1,len(nums)-1 ; j<k ;{
            // fmt.Println(j,k)
            if(nums[j]+nums[k]+nums[i]==0){
                if _,ok :=mymap[sprint([]int{nums[i],nums[j],nums[k]})]; !ok{
                    ret = append(ret,[]int{nums[i],nums[j],nums[k]})   
                    mymap[sprint([]int{nums[i],nums[j],nums[k]})] = true
                }
                k-=1
                j+=1
            }else if(nums[j]+nums[k]+nums[i]>0){
                k -=1
            }else{
                j+=1
            }
            
        }
    }
    return ret
}
