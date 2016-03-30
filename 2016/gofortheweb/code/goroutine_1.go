package main

import "fmt"
import "time"

func main() {

  sendMessage( "Hello World!")
  sendMessage( "I" )
  sendMessage( "am" )
  sendMessage( "a" )
  sendMessage( "function!" )

}

func sendMessage(n string) {
  time.Sleep( 1 * time.Second )
  fmt.Println( n )
}