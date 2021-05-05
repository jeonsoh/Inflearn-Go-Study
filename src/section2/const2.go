package main

import "fmt"

func main(){
  const a, b int = 99, 100
  const c, d, e = true, 0.84, "test1"
  const (
    x, y int16 = 59, 90
    i, k = "data", 7776
  )

  fmt.Println(a, b, c, d, e)
  fmt.Println(x, y, i ,k)
}
