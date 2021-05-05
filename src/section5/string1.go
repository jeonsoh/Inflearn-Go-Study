package main
import(
  "fmt"
  "unicode/utf8"
)
func main()  {
  //문자열
  //큰 따옴표 "", 백스쿼트 ``
  //golang: 문자 char 타입 존재하지 않음 -> rune(int32)로 문자 코드 값으로 표현
  //문자 : ''로 작성
  //자주 사용하는 escape : \\, \', \", \a (콘솔벨), \b (백스페이스), \f (쪽바꿈), \n (줄바꿈), \r (복귀), \t (탭) ...

  var str1 string = "c:\\go_study\\src\\" // -> c:\go_study\src\
  str2 := `c:\go_study\src\`

  fmt.Println(str1)
  fmt.Println(str2)

  var str3 string = "Hello, World!"
  var str4 string = "안녕하세요."
  var str5 string = "\ud55c\uae00"

  fmt.Println(str3)
  fmt.Println(str4)
  fmt.Println(str5)

  //길이 (바이트 수)
  fmt.Println(len(str3))
  fmt.Println(len(str4))

  //길이(실제 길이)
  fmt.Println(utf8.RuneCountInString(str3))
  fmt.Println(utf8.RuneCountInString(str4))
  fmt.Println(len([]rune(str4)))

}
