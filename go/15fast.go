func threeSum(nums []int) [][]int {
    
    sort.Ints(nums)
    var ret [][]int
    n := len(nums)
    for i:=0 ; i<n-2 ; i++{
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        
        low, high, sum := i+1, n-1, 0 - nums[i]
        for low<high{
            if(nums[low]+nums[high]==sum){
                ret = append(ret, []int{nums[low],nums[high],nums[i]})
                for low<high{
                    if (nums[low]==nums[low+1]){
                        low++
                    }else{
                        break
                    }
                }
                for low<high{
                    if (nums[high]==nums[high-1]){
                        high--
                    }else{
                        break
                    }
                }
                low++
                high--
            }else if(nums[low]+nums[high]<sum){
                low++
            }else{
                high--
            }
        }
    }
    return ret
}
