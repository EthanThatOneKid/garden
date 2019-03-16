package main

import (
  "fmt"
  plant "./plant"
)

func makeRules() [][]string {
  a := make([][]string, 2)
  for i := range a {
    a[i] = make([]string, 2)
  }
  a[0][0] = "A"
  a[0][1] = "AB"
  a[1][0] = "B"
  a[1][1] = "A"
  return a
}

func main() {

  rules := makeRules()
  flower := plant.Plant{"flower", []string{"A"}, rules}

  flower.Generations = flower.Grow()
  fmt.Println(flower.Generations)

}
