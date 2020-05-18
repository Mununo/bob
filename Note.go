package note
//链表定义
type LNode struct{
	Data interface{}
	Next *LNode
}
//
//创建表链
func CreateNode (node *LNode, max int){
	cur:= node
	for i:=1;i<=max; i++{
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}
//
// 打印表链的方法
func PrintNote ( info string,node *LNode){
	fmt.Print(info)
	for cur:=node.Next;cur!=nil;cur = cur.Next{
		fmt.Print(cur.Data," ")
	}
	fmt.Println()
}
