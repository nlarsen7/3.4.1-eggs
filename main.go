package main

import (
  "fmt"
  "math/rand"
  "time"
)
func dozens() {
rand.Seed(time.Now().UnixNano())
eggs:=(rand.Intn(12)+1)
fmt.Println("There are",eggs,"amount of eggs in this dozen")
}
func main() {
  dozens()
}