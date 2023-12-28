package lcof 

/*
	之所以称之为回溯算法，是因为该算法再搜索解空间时会采用“尝试”与“回退”的策略。当算法在搜索过程中遇到某个状态无法继续前进或
	无法得到满足条件的解时，它会撤销上一步的选择，退回到之前的状态，并尝试其他可能的选择
*/

//复杂的回溯问题通常包含一个或多个约束条件，约束条件通常可用于“剪枝”


func re12One(nums [][]byte,word string)bool{
	lin := len(nums)
	row := len(nums[0])

	var dfs func( i,j,k int) bool

	dfs = func( i,j,k int) bool{
		if k==len(word){
			return true
		}
		if i<0 || i>=lin ||j<0 || j>=row || nums[i][j]!=word[k]{
			return false
		}

		nums[i][j] = ' '
		dair := []int{-1,0,1,0,-1}
		jzlj := false
		for l:=0;l<4;l++{
			jzlj  = jzlj || dfs(i+dair[l],j+dair[l],k+1)
		}
		nums[i][j] = word[k]

		return jzlj
	}

	for i:=0;i<lin;i++{
		for j:=0;j<row;j++{
			if dfs(i,j,0)==true{
				return true
			}
		}
	}

	return false

	

}