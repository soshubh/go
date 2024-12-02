/*Maps are used to store data values in key:value pairs.
A map is an unordered and changeable collection that does not allow duplicates.
The default value of a map is nil.
Maps hold references to an underlying hash table.*/

// var x = map[KeyType]ValueType{key1:value1,key2:value2,...}
// x := map[KeyType]ValueType{key1:value1,key2:value2,...}

// var x = make(map[KeyType]ValueType)
// x := make(map[KeyType]ValueType)

package main
import "fmt"

func main(){
    
    var x = map[string]int{"shuhb":1,"shubh2":2,"shubh3":3}
    
    fmt.Printf("%v",x)
    
    y := make(map[string]int)
    y["shubh1"]=1
    y["shubh2"]=2
    y["shubh3"]=3
    fmt.Printf("%v",y)
}
