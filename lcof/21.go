package lcof

func re21One(nums []int) []int{
	i,j := 0,len(nums)-1

	for i<j{
		for i<j && nums[j]%2==0{
			j--
		}
		for i<j && nums[i]%2==1{
			i++
		}
		nums[j],nums[i] = nums[i],nums[j]
	}
	return nums

}