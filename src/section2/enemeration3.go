package main

import "fmt"

func main() {
  const (
    _ = iota
    A
    B
    C
  )
  const (
    _ = iota + 0.75 * 2
    DEFAULT
    SILVER
    _
    PLATIUM
  )
  fmt.Println(A, B, C)
  fmt.Println(DEFAULT, SILVER, PLATIUM)
}
