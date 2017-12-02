package main

import (
  "fmt"
  "os"

  "github.com/perfect5th/wesley/serve"
)

func startServing(args []string) {
  fmt.Println("Serving at localhost:8081")
  serve.Start(":8081")
}

var commands = map[string]func([]string){
  "serve": startServing,
}

func main() {
  var command string

  if len(os.Args) > 1 {
    command = os.Args[1]
  } else {
    fmt.Println("No command specified")
    return
  }

  commandFunc, ok := commands[command]
  if ok == true {
    if len(os.Args) > 2 {
      commandFunc(os.Args[2:])
    } else {
      commandFunc(make([]string, 0))
    }
  } else {
    fmt.Printf("Unrecognized command: %s\n", command)
  }
}
