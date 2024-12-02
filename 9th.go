package main
import "fmt"

//Slices are similar to arrays, but are more powerful and flexible.
func main(){
    //slice_name := []datatype{values}
    slice := []int{}
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
    fmt.Println(slice)
    
    temp_slice := []int{1,2,4,5,6}
    fmt.Println(len(temp_slice))
    fmt.Println(cap(temp_slice))
    fmt.Println(temp_slice)   
    
    create_slice_from_array()
}
func create_slice_from_array(){
    fmt.Print("This is Slice from array\n")
    array := [10]int{1,2,3,4,5,6,7,8,9,10}
    slice := array[2:6]
    fmt.Println(slice)
    fmt.Printf("Length  : %d\n",len(slice))
    fmt.Printf("Capacity  : %d ",cap(slice))
}
