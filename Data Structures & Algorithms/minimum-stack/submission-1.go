// Idea: use two stacks of the same height.
//   values: the normal stack
//   mins:   mins[i] = the smallest value among values[0..i]
// When we push, the new minimum is the smaller of (new value, old minimum).
// When we pop, we pop both stacks, so mins always matches values.
// Every operation is O(1).

type MinStack struct {
	values []int
	mins   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	min := val
	// if the stack is not empty, compare with the minimum so far
	if len(this.mins) > 0 && this.mins[len(this.mins)-1] < min {
		min = this.mins[len(this.mins)-1]
	}
	this.values = append(this.values, val)
	this.mins = append(this.mins, min)
}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
