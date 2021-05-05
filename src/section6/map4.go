//자료형 : 맵(4)
package main
import "fmt"
func main() {
  //맵(Map)
  //맵 조회 할 경우에 주의 할 점

  //예제1
  map1 := map[string]int{
    "apple" : 15,
    "banana": 115,
    "orange": 1115,
    "lemon": 0,
  }
  value1, ok1 := map1["lemon"]   //0 값 리턴
  value2 := map1["kiwi"]    //키가 없어서 0 리턴..
  value3, ok3 := map1["kiwi"]  //두번째 값으로 키 존재 여부 확인

  fmt.Println("ex1: ", value1, ok1)
  fmt.Println("ex1: ", value2)
  fmt.Println("ex1: ", value3, ok3)

  //예제2
  if value, ok := map1["kiwi"]; ok {
    fmt.Println("ex2: ", value)
  }else {
    fmt.Println("ex2: , kiwi is not exist!")
  }
  if value, ok := map1["banana"]; ok {
    fmt.Println("ex2: ", value)
  }else {
    fmt.Println("ex2: , banana is not exist!")
  }

  if _, ok := map1["kiwi"]; !ok{
    fmt.Println("ex2: , kiwi is not exist!")
  }

}
