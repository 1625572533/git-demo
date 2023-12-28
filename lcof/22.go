package lcof 

func re22One(heap *linked,target int) linked{

	var dfs func(link *linked, res *[]*linked)
	dfs = func(link *linked, res *[]*linked){
		if link==nil{
			return
		}
		dfs(link.next,res)
		*res = append(*res,link)
	}
	res := make([]*linked,0)
	dfs(heap,&res)

	return *(res[target-1]);
	
}

func getKthFromEnd(head *linked,k int)*linked{
	fast,slow := head,head
	for ;k>0;k--{
		fast = fast.next
	}
	for fast!=nil{
		fast = fast.next
		slow = slow.next
	}

	return slow
}