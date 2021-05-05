package main
import "fmt"
func main() {
  var a int = 50
  b := 70
  if a >= 65 {
    fmt.Println("65이상")
  }else {
    fmt.Println("65미만")
  }

  //에러 발생
    if a >= 65 {
      fmt.Println("65이상")
    }else
    {
      fmt.Println("65미만")
    }
}
