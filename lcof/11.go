package lcof 

func Re11One(n int)int{
	if n==1 || n==2{
		return n
	}
	return Re11One(n-2) + Re11One(n-1)
}