package main

import "fmt"

func main() {
  const (
    A = iota
    B
    C
  )
  const (
    Jan = iota + 1
    Feb
    Mar
  )
  fmt.Println(A, B, C)
  fmt.Println(Jan, Feb, Mar)
}
