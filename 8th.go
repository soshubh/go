package main
import "fmt"

func main(){
    // var array = [length]datatype{values}
    // array := [...]datatype{values}
    var array_1 = [3]int{1,2,3}
    var array_2 = [...]int{1,5,4,3}
    array_3 := [2]string{"Shubh","singh"}
    array_4 := [...]string{"Shubh"}
    fmt.Println(array_1,array_2,array_3,array_4)
    
    fmt.Println(array_1[1])
    array_1[1]=20
}
