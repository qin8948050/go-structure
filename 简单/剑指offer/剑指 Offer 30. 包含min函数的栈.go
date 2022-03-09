/* 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

 

示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2. */

type MinStack struct {
	Stack []int
	MinNum []int  //最小元素,辅助栈
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
	}
}

func (this *MinStack) Push(x int)  {
	this.Stack=append(this.Stack,x)
	if len(this.MinNum)==0 || this.MinNum[len(this.MinNum)-1]>=x {
		this.MinNum=append(this.MinNum, x)
	}
}


func (this *MinStack) Pop()  {
	a:=this.Stack[len(this.Stack)-1]
	this.Stack=this.Stack[:len(this.Stack)-1]
	if a == this.MinNum[len(this.MinNum)-1]{
		this.MinNum=this.MinNum[:len(this.MinNum)-1]
	}
}


func (this *MinStack) Top() int {
	var result int
	if len(this.Stack)>0{
		result=this.Stack[len(this.Stack)-1]
	}
	return result
}


func (this *MinStack) Min() int {
	var a int
	if len(this.MinNum)>0{
		a=this.MinNum[len(this.MinNum)-1]
	}
	return a
}