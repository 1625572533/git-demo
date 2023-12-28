package lcof 

func re13One(nums [][]int,target int) int{
	lin := len(nums)-1
	row := len(nums[0])-1

	var dfs func(i,j int)int 
	selected := make([][]bool,lin)
	for i:=0;i<lin;i++{
		selected[i] = make([]bool,row)
	}

	dfs =func(i,j int)int {
		if i<=lin && j<=row && (i%10 +i/10 +j/10 + j%10)<=target && !selected[i][j]{
			selected[i][j] = true
			return 1 + dfs(i+1,j) + dfs(i,j+1)
		}
		return 0
	}

	return dfs(0,0)

}