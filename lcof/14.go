package lcof

import "math"

func re14One(n int) int {
	m := n / 3
	k := m % 3
	max := 0
	if k == 0 {
		max = int(math.Pow(float64(3),float64(m)))
	}else if k==1{
		max = int(math.Pow(float64(3),float64(m-1)))* 2 *2
	}else if k==2{
		max = int(math.Pow(float64(3),float64(m-1)))* 2 
	}
	return max
}


/*
	总结出动态规划的常用术语
		将数组dp称为dp表，dp[i]表示状态i对应子问题的解
		将最小子问题对应的状态（即第1和2阶楼梯）称为初始状态
		将递推公式dp[i] = dp[i-1] + dp[i-2] 称为[状态转移方程]
*/

func cuttingRope(n int)int{

	dp := make([]int,n+1)
	dp[1] = 1

	for i:=2;i<=n;i++{
		for j:=1;j<i;j++{
			dp[i] = int(math.Max(float64(math.Max(float64(dp[i]),float64(dp[i-j]*j))),float64((i-j)*j)))
		}
	}

	return dp[10]
}