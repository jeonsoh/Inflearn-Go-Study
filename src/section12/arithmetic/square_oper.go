// 두 숫자의 연산 계산 제공 패키지(2)
package arithmetic

//x, y 제곱의 합을 리턴
func (o *Numbers) SquarePlus() int {
  return (o.x * o.x) + (o.y * o.y)
}

//x, y 제곱의 차를 리턴
func (o *Numbers) SquareMinus() int {
  return (o.x * o.x) - (o.y * o.y)
}
