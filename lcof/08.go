package lcof

type twoStack struct {
	stk1 []int
	stk2 []int
}

func (t *twoStack) appendTail(num int) {
	t.stk1 = append(t.stk1, num)
}

func (t *twoStack) deleteHead(num int) int {
	if len(t.stk2)==0 {
		res := make([]int, len(t.stk1))
		copy(t.stk1, res)
		t.stk2 = res
	}

	resp := t.stk2[len(t.stk2)]
	t.stk2 = t.stk2[:len(t.stk2)-1]
	return resp
	

}
