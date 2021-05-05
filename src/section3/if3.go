package main
import "fmt"
func main(){
  i := 100
  if i >= 120 {
    fmt.Println("120이상")
  }else if i >= 100 && i < 120 {
    fmt.Println("100이상")
  }else {
    fmt.Println("100미만")
  }
}
