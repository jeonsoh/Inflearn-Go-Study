//함수 Defer(4)
package main
import(
  "fmt"
)

func start(t string) string {
  fmt.Println("start : ", t)
  return t
}

func end(t string) {
  fmt.Println("end: ", t)
}

func a() {
  defer end(start("b")) //바로 뒤 end함수만 defer 적용, start함수는 그냥 실행  //중첩 함수 주의!
  fmt.Println("in a")
}

func main() {
  //예제1
  a()
}
