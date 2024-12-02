package main
import "fmt"

func main(){
    num := 1.553
    text := "Shubh Singh"
    
    fmt.Printf("%T",num, "%T",text) //Type of the Value
    fmt.Print("\n");
    fmt.Printf("%v" , num, "%v", text) //Value in default format
    fmt.Print("\n");
    fmt.Printf("%#v" , num, "%#v", text) //Value in go-syntex format
    fmt.Print("\n");
    fmt.Printf("%v%%" , num, "%v%%", text) // % sign
    fmt.Print("\n");
    
    formatting();
}
func formatting(){
    var temp int = 7
    fmt.Printf("%b \n",temp) //base 2
    fmt.Printf("%d \n",temp) //base 10
    fmt.Printf("%+d \n",temp) //base 10 & show sign
    fmt.Printf("%o \n",temp) //base 8
    fmt.Printf("%O \n",temp) //base 8, with leading 0o
    fmt.Printf("%x \n",temp) //base 16, lowercase
    fmt.Printf("%X \n",temp) //base 16, uppercase
    fmt.Printf("%#x \n",temp) //base 16, with leading 0x
    fmt.Printf("%4d \n",temp) //pad with spaces(width 4,right justified)
    fmt.Printf("%-4d \n",temp) //pad with spaces(width 4,right justified)
    fmt.Printf("%05d \n",temp) //pad with zeroes(make the 5 digit)
}
