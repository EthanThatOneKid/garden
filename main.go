package main

import (
  "fmt"
  // ui "github.com/manifoldco/promptui"
  plant "./plant"
)

func main() {

  // flower := plant.Plant{
  //   Species: "Pine",
  //   Axiom: "a",
  //   GrowthConfigX: []string{"a", "b", "d", "e"},
  //   GrowthConfigY: []string{"bcdedcb", "cc", "ee", ""},
  //   RenderConfigX: []string{"a", "b", "c", "d", "e"},
  //   RenderConfigY: []string{"||", "|", "|", "|", "|"},
  //   PhaseConfigX: []string{},
  //   PhaseConfigY: []string{},
  //   Gens: []string{},
  // }
  // flower.LoadGens("")
  // flower.Grow(4)

  flower := plant.Plant{
    Species: "Algae",
    Axiom: "A",
    GrowthConfigX: []string{"A", "B"},
    GrowthConfigY: []string{"AB", "A"},
    RenderConfigX: []string{"A", "B"},
    RenderConfigY: []string{"\\", "/"},
    PhaseConfigX: []string{},
    PhaseConfigY: []string{},
    Gens: []string{},
  }
  flower.LoadGens()
  flower.Grow(7)

  // flower := plant.Plant{
  //   Name: "flower",
  //   Gens: []string{"a"},
  //   GrowthConfigX: []string{"a", "b", "c"},
  //   GrowthConfigY: []string{"bac", "sb", "cs"},
  // }
  // flower.Grow(5)

  fmt.Println(flower.Gens)
  fmt.Println(flower.Render())
  // fmt.Println(flower.Print([]string{"a", "b", "c", "s"}, []string{"|", "\\", "/", " "}))

  // prompt := ui.Prompt{
	// 	Label:    "Number",
	// 	// Validate: validate,
	// }
  // res, _ := prompt.Run()
  // fmt.Println(res)

}
