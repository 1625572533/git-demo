package lcof 


func re04One(nums [][]int,taget int) bool{

	lin := len(nums)
	row := len(nums[0])

	for i:=0;i<lin;i++{

		m:= erfen(nums[i],taget)

		if m<row && nums[i][m]==taget{
			return true
		}
	}
	return false

	
}

func erfen(nums []int,taget int)int{
	i,j :=0,len(nums)-1

	for i<=j{
		m := i + (j-i)/2
		if taget >nums[m]{
			i = m+1
		}else if taget < nums[m]{
			j = m-1
		}else{
			return m
		}
	}
	return i

}

func zs04(nums [][]int,taget int)bool{

	lin := len(nums)
	row := len(nums[0])

	i,j := lin-1,0
	for i>=0 && j<row{
		if nums[i][j]  > taget{
			i--
		}else if nums[i][j]  < taget{
			j++
		}else{
			return true
		}
	}
	return false

}