package lcof 

func re10One(n int) int{
	if n==0 || n==1{
		return n
	}

	a,b := 0,1
	for i:=3;i<n;i++{
		a,b = b, a+b
	}

	return b
}


func re10DG(n int)int{
	if n==0 || n==1 {
		return n
	}

	return re10DG(n-2) + re10DG(n-1)
}