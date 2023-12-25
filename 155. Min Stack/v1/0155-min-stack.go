type MinStack struct {
	min   int
	stack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, min: 2<<31 - 1}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.min > val {
		this.min = val
	}
}
func (this *MinStack) Pop() {
	if this.min == this.stack[len(this.stack)-1] {
		// maintain minimum
		this.min = 2<<31 - 1
		for _, v := range this.stack[:len(this.stack)-1] {
			if this.min > v {
				this.min = v
			}
		}
	}
	this.stack = this.stack[:len(this.stack)-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */