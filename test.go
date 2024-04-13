package main

import (
	"bmstu/cc2024/lab1/writer"
	"os"
	"strconv"
)

func main() {
  args := os.Args
  var testnum int

  if len(args) > 1{
    testnum, _ = strconv.Atoi(args[1])
  }
	switch testnum {
	case 1:
    test1()
  case 2:
    test2()
  default:
    test1()
  }

}

func test1() {
	regex := "abb*(a|b)"
	writer.Process(regex, "abb", "aabb", "abba")
}

func test2() {
	regex := "(a|b)*abb"
	writer.Process(regex, "abb", "aabb", "babb", "babbb")
}
