package lcof

import "strings"

func re05one(nums string, s string) string {

	return strings.Replace(nums," ",s,-1)
}

func tr05Two(nums string, s string) string {

	str := strings.Builder{}
	for _,i := range nums{
		if i == ' ' {
			str.WriteString(s)
		}else{
			str.WriteRune(i)
		}
	}
	return str.String()
}