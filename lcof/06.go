package lcof 


type linked struct{
	val int
	next *linked
}
func re06One(nums *linked) []int{

	res := make([]int,0)
	for nums != nil{
		res = append(res,nums.val)
		nums = (*nums).next
	}
	n:=len(res)-1
	for i,j := 0,n;i<j;i,j = i+1,j-1{
		res[i],res[j] = res[j],res[i]
	}
	return res
}

func re06DG(nums *linked,res *[]int) {
	if nums == nil {
		return 
	}

	re06DG((*nums).next,res)
	*res = append(*res,nums.val)
}