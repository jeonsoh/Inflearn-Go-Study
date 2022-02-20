//두 숫자의 사칙연산 계산 제공 패키지(1)
package arithmetic

//x, y 두개의 Integer 구조체
type Numbers struct {
  x int
  y int
}

//x, y 합을 계산해서 반환
func (o *Numbers) Plus() int {
  return o.x + o.y
}

//x, y 차를 계산해서 반환
func (o *Numbers) Minus() int {
  return o.x - o.y
}

//x, y 곱을 계산해서 반환
func (o *Numbers) Multi() int {
  return o.x * o.y
}

//x, y 나누기를 계산해서 반환
func (o *Numbers) Divide() int {
  return o.x / o.y
}
