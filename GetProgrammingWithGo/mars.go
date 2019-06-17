package main
import "fmt"
/*
// A less clean implementation
func main(){
    fmt.Print("My weight on the Surface of Mars is ")
    fmt.Print(149.0 * 0.3783)
    fmt.Print(" lbs,  and I would be ")
    fmt.Print(41 * 365/687)
    fmt.Print(" years old.")
}
*/
func main(){
    // Printf allows you you to insert values anywhere in the text 
    fmt.Printf("My weight on the surface of Mars ins %v lbs.", 149.0*0.3783)
    fmt.Printf(" and I would be %v years old\n", 41*365/687)
}

