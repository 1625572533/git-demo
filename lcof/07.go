package lcof

type TreeNode struct {
	Val    any       // 节点值
	Height int       // 节点高度
	Left   *TreeNode // 左子节点引用
	Right  *TreeNode // 右子节点引用
}
func re07One(pre []int,in []int) *TreeNode{

	inMap := make(map[int]int,len(in))

	for i:=0;i<len(in);i++{
		inMap[in[i]] = i
	}



	return re07Xu(&pre,&inMap,0,0,len(in)-1) 
}


func re07Xu(pre *[]int,inMap *map[int]int ,i, l,r int ) *TreeNode{
	if r<l{
		return nil
	}

	root := TreeNode{Val:(*pre)[i]}
	m := (*inMap)[i]
	root.Left = re07Xu(pre,inMap,i+1,l,m-1)
	root.Right = re07Xu(pre,inMap,i+1+(m-l),m+1,r)

	return &root
}