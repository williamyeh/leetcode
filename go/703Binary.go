type KthLargest struct {
    k int
    nums []int
}


func Constructor(k int, nums []int) KthLargest {
    sort.Ints(nums)
    return KthLargest{
        k: k,
        nums: nums,
    }
}


func (this *KthLargest) Add(val int) int {
    pos := this.binary(val)
    this.nums = append(this.nums[:pos], append([]int{val},this.nums[pos:]...)...)
    // fmt.Println(pos , this.nums)
    return this.nums[len(this.nums)-this.k]
} 

func (this *KthLargest) binary(key int) int{
    l := 0
    r := len(this.nums) - 1
    
    for l <= r {
        median := (l+r)/2
        if key < this.nums[median]{
            r = median - 1
        }else if key > this.nums[median]{
            l = median + 1
        }else {
            return median
        }
    }
    return l
}

