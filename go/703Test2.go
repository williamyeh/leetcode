// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"fmt"
        "container/heap"
)

// An IntHeap is a min-heap of ints.

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

func (this *KthLargest) Pop() int {
    x := this.nums[len(this.nums)-1]
    this.nums = this.nums[:len(this.nums)-1]
    /*old := (*this).nums
    n := len(old)
    x := old[n-1]
    (*this).nums = old[0:n-1]*/
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
func main() {
 /*
  temp := []int{3,2,5}
  fmt.Println(len(temp))
  obj := Constructor(2,temp)
  fmt.Println(obj)
  fmt.Println(obj.Add(4))
  fmt.Println(obj) 
  */
temp:=[]int{4,5,8,2}
kthLargest :=  Constructor(3, temp)
fmt.Println(kthLargest.Add(3))   // return 4
fmt.Println(kthLargest.Add(5))   // return 5
fmt.Println(kthLargest.Add(10))  // return 5
fmt.Println(kthLargest.Add(9))   // return 8
fmt.Println(kthLargest.Add(4))

}
