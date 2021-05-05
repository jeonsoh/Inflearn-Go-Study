//데이터타입: BooL (1)

package main
import "fmt"
func main() {
  //BooL(Boolean): 참과 거짓
  //조건부 논리 연산자랑 주로 사용 : !, ||(or), &&(and)
  //암묵적 형 변환 X : , Nil -> False 변환 없음

  //예제1
  var b1 bool = true
  var b2 bool = false
  b3 := true
  //b4 := 1
  fmt.Println("ex1: ",b1)
  fmt.Println("ex2: ",b2)
  fmt.Println("ex3: ",b3)

  fmt.Println("ex4: ",b3 == b3)
  fmt.Println("ex5: ",b1 && b3)
  fmt.Println("ex6: ", !b1)

//error
/*
  if b4 {
    fmt.Println("ex8: true!")
  }
*/
}
