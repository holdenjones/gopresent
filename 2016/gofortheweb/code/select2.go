package main

import "fmt"
import "time"
import "math/rand"

func main() {
  c1 := randomDelay("process1")
  c2 := randomDelay("process2")
  c3 := randomDelay("process3")
  
  for i := 0 ; i < 6 ; i++ {
  
    select {
      case p1 := <-c1:
        fmt.Println( p1 )
      case p2 := <-c2:
        fmt.Println( p2 )
      case p3 := <-c3:
        fmt.Println( p3 )
    }
    
  }

}

func randomDelay(n string) (chan string) {
  c := make(chan string)
  go func() {
    for {
      time.Sleep( time.Duration(rand.Intn(5)) * time.Second )
      c <- n
    }
  }()
  
  return c
}