package lcof

func re23One(head *linked) *linked {
	fast := head
	slow := &linked{}


	for fast!= nil {
		next := fast.next
		fast.next = slow.next
		slow.next = fast
		fast = next

	}
	return slow.next
}

func re23DG(head *linked) *linked {
	if (head==nil || head.next==nil){
		return head
	}
	ans := re23DG(head.next)
	head.next.next = head 
	head.next = nil 
	return ans 
}
