//자료형: 배열(2)
package main
import "fmt"
func main() {
  //배열 순회

  //예제1
  arr1 := [5]int{1, 10, 100, 1000, 10000}
  //len 길이 반복
  for i := 0; i < len(arr1); i++ {
    fmt.Println("ex1: ", arr1[i])
  }
  fmt.Println()

  //예제2
  arr2 := [5]int{7, 77, 777, 7777, 77777}
  for i, v := range arr2 {
    fmt.Println("ex2: ", i, v)
  }
}
