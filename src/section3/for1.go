package main

import "fmt"

func main() {
  //반복문 - for
  //Go에서 유일하게 제공되는 반복문
  //다양한 사용법 숙지

  //예제1
  for i := 0; i < 5; i++ {
    fmt.Println("ex1: ", i)
  }

  //예제2 무한루프
  // for {
  //   fmt.Println("hello, golang!")
  // }

  //예제3 Range 용법
  loc := []string{"Seoul", "Busan", "Incheon"}
  for index, name := range loc {
    fmt.Println("ex3: ", index, name)
  }
  for _, name := range loc {
    fmt.Println("ex3: ", name)
  }

  //에러 발생 - 괄호 위치
  // for i := 0; i < 5; i++
  // {
  //
  // }

  //에러 발생
  // for i := 0; i < 5; i++
  // fmt.Println("ex1: ", i)



}
