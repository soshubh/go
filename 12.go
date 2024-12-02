package main
import "fmt"

func add(x int,y int) int{
    return x+y
}
func namedReturnValue(x int,y int)(z int){
    z=x+y
    return
    //return z
}
func addStringReturn(x int, y int,s string)(z int, p string){
    z=x+y
    p=s+" Singh"
    return
}

func main(){
    x :=3
    y :=5
    fmt.Println(add(x,y))
    
    fmt.Println(namedReturnValue(x,y))
    
    fmt.Println(addStringReturn(x,y,"shubh"))
    
    //Store return value
    _,p := addStringReturn(x,y,"shubh")
    fmt.Println(p)
}
