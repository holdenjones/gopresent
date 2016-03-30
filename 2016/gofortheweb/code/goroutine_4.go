package main

import "fmt"
import "time"
import "math/rand"

func main() {
  c := make(chan int)

  go handleUser("process1",c)
  go handleUser("process2",c)
  go handleUser("process3",c)
  
  for i := 0 ; i < 6 ; i++ {
    time.Sleep( 1 * time.Millisecond )
    c <- i
  }
  
  time.Sleep( 30 * time.Second )

}

func handleUser(n string , in chan int) {
  for {
    userInt := <-in
    time.Sleep( time.Duration(rand.Intn(5)) * time.Second )
    fmt.Println( n , " finished: [" , userInt , "]" ) 
  }
}