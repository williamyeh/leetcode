type KthLargest struct {
    k int
    nums []int
}

func (this KthLargest) Len() int {
    return len(this.nums)
}

func (this KthLargest) Less(i,j int) bool {
    return this.nums[i] < this.nums[j]
}

func (this KthLargest) Swap(i, j int) {
    this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

func (this *KthLargest) Push(x interface{}) {
	this.nums = append(this.nums, x.(int))
}

func (this *KthLargest) Pop() interface{} {
	x := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return x
}

func (this *KthLargest) Add(val int) int {
    heap.Push(this, val)
    if len(this.nums) > this.k{
        heap.Pop(this)
    }
    return this.nums[0]
}

func Constructor(k int, nums []int) KthLargest {
    obj := KthLargest{
        k: k,
        nums: make([]int, 0 ,k+1),
    }
    
    for _, num:= range nums{
        obj.Add(num)
    }
    return obj
}

