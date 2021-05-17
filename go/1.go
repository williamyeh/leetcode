func twoSum(nums []int, target int) []int {
    
    mymap := map[int]int{}
    
    for i,j := range(nums){
        if v, ok := mymap[j]; ok {
            return []int{i,v}
        }
        mymap[target-j]=i
    }
    return nil
}
