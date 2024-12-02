package main
import "fmt"

//Slices are similar to arrays, but are more powerful and flexible.
func main(){
    //slice_name := []datatype{values}
    slice := []int{}
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
    fmt.Println(slice)
    
    //Append element
    //slice_name = append(slice_name, element1, element2, ...)
    slice = append(slice, 4,3,2)
    fmt.Print(slice,"\n")
    fmt.Printf("Cap :%d\nLen :%d",cap(slice),len(slice))
    
    //Append One Slice To Another 
    //slice3 = append(slice1, slice2...)
}
