package main
import (
  "fmt"
  "strings"
  )
func main() {
  //결합 : 일반 연산
  str1 := "The Go programming language is an open source project to make programmers more productive." +
          "Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that" +
          "get the most out of multicore and networked machines, while its novel type system enables flexible and modular"

  str2 := "A brief Hello, World tutorial to get started. Learn a bit about Go code, tools, packages, and modules."

  fmt.Println(str1 + str2) //string은 수정이 불가능, 이렇게 쓰면 계속해서 new하게 됨. 때문에 다른 언어에서는 이 방법 대신 buffer를 사용하곤 함.

  //결합 : Join
  strSet := []string{} //슬라이스 선언
  strSet = append(strSet, str1) //내부적으로 연산을 최소화 할 수 있다.
  strSet = append(strSet, str2)

  fmt.Println("ex2: ", strings.Join(strSet, "----")) //두 문장이 결합되는 중간 부분을 ----로 구분하겠다.
}
