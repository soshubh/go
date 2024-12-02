package main
import ("fmt")

func main(){
    var first string
    first ="Shubh"
    fmt.Print(first+"\n")
    fun1()
}

// temp := 333 -> Can only be used inside functions

var second int
func fun1(){
    second = 884
    fmt.Println(second)
    var a,b,c = 20,true,"shubh"
    fmt.Print(a,b,c)
}
