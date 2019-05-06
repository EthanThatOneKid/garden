package main

// +--------------+
// | Dependencies |
// +--------------+

import (
  "fmt"
  "strings"
  // ui "github.com/manifoldco/promptui"
  plant "./plant"
  static "./static"
)

// +---------+
// | Helpers |
// +---------+

func restorePlant(plants []string, index int) (Plant) {
  i := i * 7
  plants := static.plants
  return plant.Plant{
    Species: string(plants[i + 0]),
    Entry: string(plants[i + 1]),
    Axiom: string(plants[i + 2]),
    GrowthConfigX: strings.Split(string(plants[i + 3]), ","),
    GrowthConfigY: strings.Split(string(plants[i + 4]), ","),
    RenderConfigX: strings.Split(string(plants[i + 5]), ","),
    RenderConfigX: strings.Split(string(plants[i + 6]), ","),
  }
}

// +--------------+
// | Main Process |
// +--------------+

func main() {

  flower := plant.Plant{
    Species: "Algae",
    Axiom: "A",
    GrowthConfigX: []string{"A", "B"},
    GrowthConfigY: []string{"AB", "A"},
    RenderConfigX: []string{"A", "B"},
    RenderConfigY: []string{"\\", "/"},
  }
  flower.LoadGens()
  flower.Grow(7)
  fmt.Println(flower.Gens)
  fmt.Println(flower.Render())
  
  // prompt := ui.Prompt{
	// 	Label:    "Number",
	// 	// Validate: validate,
	// }
  // res, _ := prompt.Run()
  // fmt.Println(res)

}
