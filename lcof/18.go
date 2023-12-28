package lcof

func re18One(link *linked, target int) *linked {
	if link.val == target {
		return link.next
	}

	head := link

	for link != nil {
		pre := link
		link := link.next
		if link.val == target {
			pre.next = link.next
			break
		}

	}

	return head
}
