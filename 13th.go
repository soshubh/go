package main
import "fmt"

func myFun(x int) int{
    if x==10{
     return 0
    }
    fmt.Println(x)
    return myFun(x+1)
}
func factorial(x int) (fact int){
    if x > 0{
        fact = x * factorial(x-1)
    }    else{
        fact=1
    }
    return
}

func main(){
    myFun(1)
    x:=factorial(5)
    fmt.Println("Factorial of 5: ",x)
    
}
