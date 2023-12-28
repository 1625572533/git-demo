package lcof 

func re15One(n uint32)int{
	res :=0

	for n!=0{
		if (n&1)==1{
			res ++
		} 
		n >>=1
	}
	return res
}

func re15Two(n uint32)int{
	res :=0

	for n!=0{
		n &= n-1
		res ++
	}
	return res
}