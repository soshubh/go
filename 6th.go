package main
import "fmt"

func main(){
    var first, second = "Shubh","Singh"
    fmt.Print(first,second+"\n")

    fmt.Println(first,second)
    
    fmt.Printf("Type- %t, Value-%v",first,first)
    fmt.Printf("Type- %t, Value-%v",second,second)
    
}
