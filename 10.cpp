package main
import "fmt"
//Memory Efficiency
func main(){
    sliceOld := []int{1,2,3,4,5,6,7,8,9,10}
    fmt.Print(sliceOld,"\n")
    fmt.Printf("Len : %d\nCap : %d",len(sliceOld),cap(sliceOld))
    
    neededNumebrs := sliceOld[:len(sliceOld)-5]
    sliceNew := make([]int,len(neededNumebrs));
    copy(sliceNew,neededNumebrs)
    fmt.Print("\n",sliceNew)
    fmt.Printf("\nLen : %d\nCap : %d",len(sliceNew),cap(sliceNew))
    
}
