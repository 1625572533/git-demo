package lcof

import "math"

func re17One(n int) []int {
	sum := int(math.Pow(10,float64(n+1)))
	res := make([]int,sum-1)

	for i,_ := range res{
		res[i] = i+1
	}
	return res

}

//用溯源算法实现
func print(n int)[]string{

	ans := make([]string,0)
	s := make([]byte,0)

	var dfs func(i,j int)
	dfs = func(i,j int){
		if i==j{
			ans = append(ans,string(s))
			return 
		}

		k:=0 
		if i==0{
			k++
		}

		for k<10{
			s = append(s,byte('0'+k))
			dfs(i+1,j)
			s = s[:len(s)-1]
			k++
		}

	}
	for i:=1;i<=n;i++{
		dfs(0,i)
	}
	return ans
}