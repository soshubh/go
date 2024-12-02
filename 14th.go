//A struct is used to create a collection of members of different data types, into a single variable.
/* type name_ struct {
    first datatype
    second datatype
    .....
}*/

package main
import "fmt"

type first_struct struct{
    myName string
    age int
    isActive bool
}
func printFun(shubh2 first_struct){
     fmt.Println(shubh2.myName,shubh2.age,shubh2.isActive)
}

func main(){
   var shubh first_struct
   shubh.myName="Shuhb"
   shubh.age=24
   shubh.isActive=true
   printFun(shubh)
}
