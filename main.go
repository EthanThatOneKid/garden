package main

import (
  "fmt"
  plant "./plant"
)

func makeRules() [][]string {


}

func main() {

  flower := plant.Plant{
    "flower",
    []string{"A"},
    [][]string{{"A", "AB"}, {"B", "A"}}
  }

  fmt.Println(flower)

}
