package main 
import "fmt"
func SlicesIndex[S ~[]E,E comparable](s S,v E) int{
	for i:=range s{
		if v==s[i]{
			return i
		}
	}
	return -1
}
type List[T any] struct{
	head,tail  *element[T]
}
type element[T any] struct{
	next *element[T]
	val T
}
func (lst *List[T]) Push(v T){
	if lst.tail==nil{
		lst.head=&element[T]{val:v}
		lst.tail=lst.head
	}else{
		lst.tail.next=&element[T]{val:v}
		lst.tail=lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T{
	var elems []T
	for e:=lst.head;e!=nil;e=e.next{
		elems=append(elems,e.val)
	}
	return elems
} 
func main(){
	var s=[]string{"foo","bar","zoo"}
	fmt.Println("index of zoo:",SlicesIndex(s,"zoo"))
	_=SlicesIndex[[]string,string](s,"zoo")
	lst:=List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:",lst.AllElements())

}