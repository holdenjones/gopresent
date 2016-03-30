package main

import "fmt"
import "time"

func main() {

  go sendMessage( "Hello World!")
  fmt.Println( "Starting 10 second timer." )
  time.Sleep( 10 * time.Second )
  fmt.Println( "Time's up!" )
}

func sendMessage(n string) {
  for i := 0 ; i < 5 ; i++ {
    time.Sleep( 1 * time.Second )
    fmt.Println( n )
  }
}