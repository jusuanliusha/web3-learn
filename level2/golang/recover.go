package main
import "fmt"

func mapPanic(){
	panic("a problem")
}
func main(){
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("Recovered. Error:\n",r)
		}
	}()
	mapPanic()
	fmt.Println("After mapFanic()")
}