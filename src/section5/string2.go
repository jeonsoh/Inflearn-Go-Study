package main
import "fmt"
func main() {
  //문자열 표현
  //문자열 : UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의!

  var str1 string = "Golang"
  var str2 string = "World"
  var str3 string = "고프로그래밍"
  //아스키 코드 값 출력
  fmt.Println("ex: ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
  fmt.Println("ex: ", str2[0], str2[1], str2[2], str2[3], str2[4])
  fmt.Println("ex: ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

  fmt.Printf("ex: %c %c %c %c %c %c \n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
  fmt.Printf("ex: %c %c %c %c %c \n", str2[0], str2[1], str2[2], str2[3], str2[4])
  fmt.Printf("ex: %c %c %c %c %c %c \n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5]) //한글 깨
  //format형에서 그냥 배열에 접근하면 한글 깨. rune형 배열 사용

  conStr := []rune(str3)
  fmt.Printf("ex: %c %c %c %c %c %c \n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5]) //한글 정상 출력

  for i, char := range str1 {
    fmt.Printf("ex3: %c(%d)\t", char, i)
  }
  fmt.Println()
  for i, char := range str2 {
    fmt.Printf("ex4: %c(%d)\t", char, i)
  }

}
