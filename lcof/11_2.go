package lcof 

func re11_2One(nums []int)int{
	for i:=0;i<len(nums)-1;i++{
		if nums[i+1] <nums[i]{
			return nums[i+1]
		}
	}
	return nums[0]
}

func re11EF(nums []int)int{
	l,r := 0,len(nums)-1
	for l<r{
		m := l + (r-l)/2
		if nums[m] > nums[r]{
			l = m+1
		}else if nums[m] < nums[r]{
			r = m
		}else{
			r--
		}
	}

	return nums[l]
}

func re11EF2(nums []int)int{
	l,r := 0,len(nums)-1
	for l<r{
		if nums[l] <nums[r]{
			break
		}
		m:= l+(r-l)/2

		if nums[m] > nums[l]{
			l = m+1
		}else if nums[m] < nums[l]{
			r= m
		}else{
			l++
		}
		
	}
	return nums[l]
}