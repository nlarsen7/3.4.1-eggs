//Nicholas Larsen
//4/25/2020
//picks a random number and outputs the random number
package main

import (
  "fmt"
  "math/rand"
  "time"
)
func dozens() {
rand.Seed(time.Now().UnixNano())
eggs:=(rand.Intn(12)+1)
//random number of eggs in a dozen
fmt.Println("There are",eggs,"amount of eggs in this dozen")
//Prints out the amount of eggs
}
func main() {
  dozens()
  //outputs what the Dozen function does
}