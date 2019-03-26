package main

import (
  "fmt"
  plant "./plant"
)

func main() {

  flower := plant.Plant{"flower", []string{"A"}, []string{"A", "B"}, []string{"AB", "A"}}

  flower.Grow(5)
  fmt.Println(flower.Gens)

}
