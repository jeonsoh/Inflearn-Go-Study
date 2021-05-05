package main

import "fmt"

func main() {
  //제어몬(조건문)
  //if문: 반드시 boolean 검사 -> 1, 0 사용 불가 : 자동 형 변환 불가
  //소괄호 사용 x
  var a int = 20
  b := 20

  if a >= 15 {
    fmt.Println("15이상")
  }
  if b >= 25 {
    fmt.Println("25이상")
  }
  if c:= 40; c >= 35 {
    fmt.Println("35이상")
  }

  //에러 발생 1
  //byte code로 만들때 문장 마지막에 ;를 붙여서 인식하기때문에 괄호가 다음 행으로 넘어가면 에러발생
  if b > 15
  {

  }

  //에러 발생 2
  //괄호 생략
  if b >= 25
    fmt.Println("25이상")

  //에러 발생 3
  if c:= 1; c {
    fmt.Println("True")
  }
}
