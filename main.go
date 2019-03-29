package main

import (
  "fmt"
  // ui "github.com/manifoldco/promptui"
  plant "./plant"
)

func main() {

  flower := plant.Plant{
    Name: "flower",
    Gens: []string{"A"},
    GrowthConfigX: []string{"A", "B"},
    GrowthConfigY: []string{"AB", "A"},
  }
  flower.Grow(5)
  // flower.Phase([]string{"A"}, []string{"F"}, 2)
  fmt.Println(flower.Gens)
  fmt.Println(flower.Print())

  // prompt := ui.Prompt{
	// 	Label:    "Number",
	// 	// Validate: validate,
	// }
  // res, _ := prompt.Run()
  // fmt.Println(res)

}
