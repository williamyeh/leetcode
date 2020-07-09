func runningSum(nums []int) []int {
    sum := 0
    for k,v := range(nums){
        sum += v
        nums[k] = sum
    }
    return nums
}
