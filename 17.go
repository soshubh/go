package main
import "fmt"

func main(){
    x := map[int]int{10:1,20:2}
    fmt.Println(x[10])
    
    x[10]=10
    fmt.Println(x[10])
    
    x[50]=5
    fmt.Println(x[50])
    
    delete(x,50)
    
    fmt.Println(x)
    
    //Check For Specific Elements in a Map
    
    tempValue1, aBool1 := x[20]
    tempValue2, aBool2 := x[30]
    _, aBool3 := x[10]
    
    fmt.Println(tempValue1,aBool1)
    fmt.Println(tempValue2,aBool2)
    fmt.Println(aBool3)
    
    
}
