//함수 Closure(1)
package main
import "fmt"
func main() {
  //클로저(closure)
  //익명함수 사용할 경우 함수를 변수에 할당해서 사용 가능
  //함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언 된 변수에 접근 가능한 함수
  //함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
  //함수를 호출할 때 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트 -> 무분별한 전역변수 남발 ...
  //남발 -> 객체들이 메모리에 존재하므로 -> 메모리 부족, 오버플로우 현상, 자원을 무뷴별하게 사용할 가능성이 있다.
  //따라서 클로저 정확하게 이해하고 사용해야 한다.

  //예제1
  multiply := func(x, y int) int { //익명 함수
    return x * y
  }

  r1 := multiply(5, 10)
  fmt.Println("ex1: ", r1)

  //예제2 closure 사용
  m, n := 5, 10 //변수가 캡처
  sum := func(c int) int { //익명함수 변수 할당
    return m + n + c //지역변수 소멸되지 않는다.(함수 호출 시 마다 사용 가능 )
  }
  r2 := sum(10)
  fmt.Println("ex2: ", r2)
}
