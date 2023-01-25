package main

import (
  "fmt"
  "os"
  "monkey/repl"
)

func main() {
  fmt.Println("This is Monkey.")
  repl.Start(os.Stdin, os.Stdout)
}