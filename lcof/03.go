package lcof

func re03(nums []int) int{
	map2 := make(map[int]struct{})

	for i:=0;i<len(nums);i++{
		if _,ok :=map2[i]; ok{
			return i
		}
		map2[i] = struct{}{}
	}

	return -1
}

func ydjh(nums []int) int{
	
	for i:=0;i<len(nums);i++{
		if nums[i]==i{
			continue
		}
		if nums[i] == nums[nums[i]]{
			return nums[i]
		}
		nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
	}
	return -1
}