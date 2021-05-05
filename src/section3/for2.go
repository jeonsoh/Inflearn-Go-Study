package main
import "fmt"
func main() {
  //ex1
  sum1 := 0
  for i:= 0; i <= 100; i++ {
    sum1 += i
  }
  fmt.Println("ex1: ", sum1)
  //ex2
  sum2, i := 0, 0
  for i <= 100 {
    sum2 += i
    i++
    // j := i++ //Go에서 후치연산은 반환 값 x
  }
  fmt.Println("ex2: ", sum2)
  //ex3
  sum3, i := 0, 0
  for {
    if i > 100 {
      break
    }
    sum3 += i
    i++
  }
  fmt.Println("ex3: ", sum3)
  //ex4
  for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
    fmt.Println("i-> ",i, "j ->", j)
  }

//에러 발생. 후치 연산은 반환값 없음
  // for i, j := 0, 0; i <= 10; i++, j++{
  //   fmt.Println("i-> ",i, "j ->", j)
  // }

}
