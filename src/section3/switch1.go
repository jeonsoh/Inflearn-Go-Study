package main
import "fmt"

func main(){
  //제어문(조건문) - switch
  //Switch 뒤 표현식(Expression) 생략 가능
  //case 뒤 표현식(Expression) 사용 가능
  //자동 break 때문에 falltorough 존재
  //Type분기 -> 값이 아닌 변수 Type으로 분기 가능

  a := -7
  switch {
  case a < 0:
    fmt.Println(a, "음수")
  case a == 0:
    fmt.Println(a, "0")
  case a > 0:
    fmt.Println(a, "양수")
  }

  switch b := 27; {
  case b < 0:
    fmt.Println(b, "음수")
  case b == 0:
    fmt.Println(b, "0")
  case b > 0:
    fmt.Println(b, "양수")
  }

  switch c := "go"; c {
  case "go":
    fmt.Println("go!")
  case "java":
    fmt.Println("java!")
  default:
    fmt.Println("일치하는 값 없음")
  }

  switch c := "go"; c + "lang" {
  case "golang":
    fmt.Println("golang!")
  case "java":
    fmt.Println("java!")
  default:
    fmt.Println("None!")
  }

  switch i, j := 20, 30; {
  case i < j:
    fmt.Println("i < j")
  case i == j:
    fmt.Println("i == j")
  case i > j:
    fmt.Println("i > j")
  }
}
