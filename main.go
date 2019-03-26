package main

import (
  "fmt"
  plant "./plant"
)

func main() {

  flower := plant.Plant{"flower", []string{"A"}, []string{"A", "B"}, []string{"AB", "A"}}
  flower.Grow(5)
  flower.Phase([]string{"A"}, []string{"F"})
  fmt.Println(flower.Gens)

}
