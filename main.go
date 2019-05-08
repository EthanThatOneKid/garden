package main

// +--------------+
// | Dependencies |
// +--------------+

import (
  "fmt"
  "strings"
  ui "github.com/manifoldco/promptui"
  . "./plant"
  static "./static"
)

// +---------+
// | Helpers |
// +---------+

func restorePlant(plants []string, index int) (Plant) {
  i := index * 7
  fmt.Println(static.Plants)
  //plants := static.Plants
  return Plant{
    Species: string(plants[i + 0]),
    Entry: string(plants[i + 1]),
    Axiom: string(plants[i + 2]),
    GrowthConfigX: strings.Split(string(plants[i + 3]), ","),
    GrowthConfigY: strings.Split(string(plants[i + 4]), ","),
    RenderConfigX: strings.Split(string(plants[i + 5]), ","),
    RenderConfigY: strings.Split(string(plants[i + 6]), ","),
  }
}

func visitGarden() {
  fmt.Println("Visit Garden")
  // flower := plant.Plant{
  //   Species: "Algae",
  //   Axiom: "A",
  //   GrowthConfigX: []string{"A", "B"},
  //   GrowthConfigY: []string{"AB", "A"},
  //   RenderConfigX: []string{"A", "B"},
  //   RenderConfigY: []string{"\\", "/"},
  // }
  // flower.LoadGens()
  // flower.Grow(7)
  // fmt.Println(flower.Gens)
  // fmt.Println(flower.Render())
}

func checkPlantDex() {
  fmt.Println("Check plant dex")
}

func viewGardenerLicense() {
  fmt.Println("viewGardenerLicense")
}

// +--------------+
// | Main Process |
// +--------------+

func main() {

  fmt.Println(static.Plants)

  for true {

    fmt.Println("(1) Visit Garden")
    fmt.Println("(2) Check Plant Dex")
    fmt.Println("(3) View Gardener License")
    prompt := ui.Prompt{"> "}
    res, _ := prompt.Run()
    switch res {
      case "1":
        visitGarden()
      case "2":
        checkPlantDex()
      case "3":
        viewGardenerLicense()
      default:
        fmt.Println("\n")
    }

  }

}
