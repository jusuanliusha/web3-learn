package main
import "os"
func main(){
	panic("a problem")
	_,err:=os.Create("/tmp/temp11/file")
	if err!=nil{
		panic(err)
	}
}