package lcof 

func mergeTwoLists(l1 *linked,l2 *linked) *linked{

	dump := &linked{}
	head := dump
	for l1!=nil && l2!=nil{
		if l1.val <=l2.val{
			dump.next = l1 
			l1 = l1.next
		}else{
			dump.next=l2
			l2 = l2.next
		}
	}

	if l2==nil{
		dump.next = l1
	}else{
		dump.next=l2
	}
	return head.next
}

func re24DG(l1 *linked,l2 *linked)*linked{
	
}