package main
import "fmt"
func main() {
  a := 30 / 15
  switch a {
  case 2, 4, 6:
    fmt.Println("a -> ", a, " 짝수")
  }

//break가 없어도 go는 자동으로 break가 되서 다른 케이스 안돈다.
  switch e := "go"; e {
  case "java":
    fmt.Println("Java")
    fallthrough
  case "go":
    fmt.Println("go")
    fallthrough //조건에 맞지 않아도 다음 케이스 진입
  case "python":
    fmt.Println("python")
    fallthrough
  case "c":
    fmt.Println("c")
    //fallthrough //맨 마지막 case문에서는 사용 불가
  }
}
