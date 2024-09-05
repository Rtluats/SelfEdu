package stack

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.



Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


Constraints:

-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 104 calls will be made to push, pop, top, and getMin.
*/

type MinStack struct {
	minElems []int
	store    []int
}

func Constructor() MinStack {
	return MinStack{
		minElems: make([]int, 0),
		store:    make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.store) == 0 {
		this.minElems = append(this.minElems, val)
	} else if val <= this.minElems[len(this.minElems)-1] {
		this.minElems = append(this.minElems, val)
	}
	this.store = append(this.store, val)
}

func (this *MinStack) Pop() {
	lst := this.store[len(this.store)-1]
	if lst == this.minElems[len(this.minElems)-1] {
		this.minElems = this.minElems[:len(this.minElems)-1]
	}
	this.store = this.store[:len(this.store)-1]
}

func (this *MinStack) Top() int {
	return this.store[len(this.store)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElems[len(this.minElems)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
