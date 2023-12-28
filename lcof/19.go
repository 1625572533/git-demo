package lcof 

func re19One(s string,p string)bool{
	m := len(s)
	n := len(p)

	dp := make([][]bool,m+1)
	for i,_ := range dp{
		dp[i] = make([]bool,n+1)
	}

	dp[0][0] = true

	for i:=0;i<=m;i++{
		for j:=1;j<=n;j++{
			if p[j-1]=='*'{
				dp[i][j] = dp[i][j-2] // 匹配0次的情况
				if i>0 && (p[j-2] == s[i-1] || p[j-2] == '.'){
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}else if i>0 && (p[j-1]=='.' || p[j-1] == s[j-1]){
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}