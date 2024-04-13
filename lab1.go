package main

import (
	"bmstu/cc2024/lab1/writer"
	"os"
)

func main() {
  args := os.Args
  
  var regexp string
  var exp string

  if len(args) > 1{
    regexp = args[1]
  }
  if len(args) > 2{
    exp = args[2]
  }

  writer.Process(regexp, exp)

}

