package main

import "fmt"
import "time"

func main() {
  c := make(chan string)

  go sendMessage( "Hello World!" , c)
  
  myMessage := <-c
  
  fmt.Println( myMessage );

}

func sendMessage(n string, c chan string) {
  time.Sleep( 1 * time.Second )
  c <- n
}