package main
import "fmt"
func main() {
  //예제1
  fmt.Println("ex1: ", true && true)
  fmt.Println("ex2: ", true && false)
  fmt.Println("ex3: ", false && false)
  fmt.Println("ex4: ", true && true)
  fmt.Println("ex5: ", true && false)
  fmt.Println("ex6: ", false && false)
  fmt.Println("ex7: ", !true)
  fmt.Println("ex8: ", !false)

  //예제2
  num1 := 15
  num2 := 37

  fmt.Println("ex2: ", num1 < num2)
  fmt.Println("ex2: ", num1 > num2)
  fmt.Println("ex2: ", num1 >= num2)
  fmt.Println("ex2: ", num1 <= num2)
  fmt.Println("ex2: ", num1 == num2)
  fmt.Println("ex2: ", num1 != num2)

}
