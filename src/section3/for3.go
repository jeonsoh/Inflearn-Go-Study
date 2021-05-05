package main
import "fmt"

func main() {
  //ex1
  Loop1:
   for i:=0; i<5; i++ {
     for j:=0; j<5; j++ {
       if i==2 && j==4 {
         break Loop1 //가장 바깥으로 빠져 나감. Loop를 없애고 break만 쓰면 가장 가까운 for만 종료
       }
       fmt.Println(i, j)
     }
   }
   //ex2
   for i:=0; i<10; i++ {
     if i % 2 == 0 {
       continue
     }
     fmt.Println(i)
   }

   //에러 발생 - Loop 레이블 밑에 관련 없는 소스코드

   //ex3
   Loop2:
  for i:=0; i<3; i++ {
    for j:=0; j<3; j++ {
      if i==1 && j==2 {
        continue Loop2
      }
      fmt.Println(i, j)
    }
  }
}
