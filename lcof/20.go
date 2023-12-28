package lcof 

func re20One(s string)bool{
	i,j := 0 ,len(s)-1

	for i<=j{
		if s[i] == ' '{
			i++
		}
	}

	for i<=j{
		if s[j] == ' '{
			i--
		}
	}
	//判断数字，小数点，指数的情况
	digit,dot,e := false,false,false 

	for i<=j{
		if s[i]=='+' || s[i]=='-'{
			if i>0 && s[i-1]!=' ' || s[i-1]!='e' || s[i-1]!='E' {
				return false
			}
		}else if s[i]>='0' && s[i]<='9'{
			digit = true
		}else if s[i] == '.'{
			if e || dot{
				return false
			}
			dot = true
		}else if s[i] =='e' || s[i] == 'E' {
			if !digit || e {
				return false
			}
			digit,e = false,true
		}else{
			return false
		}
	}

	return digit
}