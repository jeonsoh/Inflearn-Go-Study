//함수 기초 (3)
package main
import "fmt"

func multiply(x int, y int) (int, int) { //(x, y int) 가능
  return x * 10, y * 10
}

func arraymultiply(a, b, c, d, e int) (int, int, int, int, int) {
  return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
  //다중 값 반환
  //예제1
  a, b := multiply(10, 5)
  // c := multiply(10, 5) //변수 개수 부족
  c, _ := multiply(10, 5)
  _, d := multiply(10, 5)

  fmt.Println("ex1: ", a, b)
  fmt.Println("ex2: ", c)
  fmt.Println("ex3: ", d)

  //예제2
  x1, x2, x3, x4, x5 := arraymultiply(1, 2, 3, 4, 5)
  y1, _, y3, _, y5 := arraymultiply(1, 2, 3, 4, 5)
  fmt.Println("ex2: ", x1, x2, x3, x4, x5)
  fmt.Println("ex2: ", y1, y3, y5)

}
