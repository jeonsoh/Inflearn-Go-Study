//구조체 심화 (1)
package main
import "fmt"

func NewAccount(number string, balance float64, interest float64) *Account { //포인터 반환 아닌 경우 값 복사
  return &Account{number, balance, interest} //구조체 인스턴스를 생성 한 뒤 포인터를 리턴

}
type Account struct {
  number string
  balance float64
  interest float64
}

func main() {
  //구조체 생성 패던 예제

  //예제1
  kim := Account{number: "245-901", balance: 1000000, interest: 0.015}

  var lee *Account = new(Account)
  lee.number = "245-902"
  lee.balance = 12000000
  lee.interest = 0.025

  fmt.Println("ex1: ",kim)
  fmt.Println("ex1: ",lee)

  //예제2
  park := NewAccount("245-903", 130000000, 0.035)
  fmt.Println("ex2: ",park)

}
