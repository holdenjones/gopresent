package main

import "fmt"
import "time"
import "math/rand"

func main() {
  c1 := randomDelay("process1")
  c2 := randomDelay("process2")
  c3 := randomDelay("process3")
  
  for i := 0 ; i < 6 ; i++ {
    fmt.Println( <-c1 )
    fmt.Println( <-c2 )
    fmt.Println( <-c3 )
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