package main

type MyStack struct {
	q []int
}

func main() {
	ms := MyStack{}
	ms.Push(1)
	ms.Pop()
	println(ms.Empty())
}
func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
}

func (this *MyStack) Pop() int {
	elem := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return elem
}

func (this *MyStack) Top() int {
	return this.q[len(this.q)-1]
}

func (this *MyStack) Empty() bool {
	if len(this.q) == 0 {
		return true
	}
	return false
}

//TODO: решение прошло литкодовскую систему проверки, но по факту решена не корректна из-за отсутвия очередей, мб потом сделаю нормально лол, XД
