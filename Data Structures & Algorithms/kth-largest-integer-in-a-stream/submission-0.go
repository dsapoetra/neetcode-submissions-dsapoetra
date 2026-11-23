// Idea: we only ever need the k biggest numbers. Keep them in a sorted
// list (smallest first). Then the k-th largest number of everything is
// simply the FIRST number of that list.
// Add: put the new number at its sorted place; if the list now has more
// than k numbers, throw away the smallest one.
// Add is O(k). (A min-heap would make it O(log k), but needs more code.)

type KthLargest struct {
	k   int
	top []int // the k biggest numbers so far, sorted, smallest first
}

func Constructor(k int, nums []int) KthLargest {
	this := KthLargest{k: k}
	for _, num := range nums {
		this.Add(num)
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	// add at the end, then move it left until the list is sorted again
	this.top = append(this.top, val)
	for i := len(this.top) - 1; i > 0 && this.top[i] < this.top[i-1]; i-- {
		this.top[i], this.top[i-1] = this.top[i-1], this.top[i]
	}

	// keep only the k biggest
	if len(this.top) > this.k {
		this.top = this.top[1:]
	}
	return this.top[0]
}
