package main

// +--------------+
// | Dependencies |
// +--------------+

import (
  "os"
  "fmt"
  "strings"
  ui "github.com/manifoldco/promptui"
  . "./static" // Plants ([]string)
  . "./plant" // Plant (struct)
)

// +---------+
// | Helpers |
// +---------+

func restorePlant(index int) (Plant) {
  i := index * 7
  return Plant{
    Species: string(Plants[i + 0]),
    Entry: string(Plants[i + 1]),
    Axiom: string(Plants[i + 2]),
    GrowthConfigX: strings.Split(string(Plants[i + 3]), ","),
    GrowthConfigY: strings.Split(string(Plants[i + 4]), ","),
    RenderConfigX: strings.Split(string(Plants[i + 5]), ","),
    RenderConfigY: strings.Split(string(Plants[i + 6]), ","),
  }
}

func visitGarden() {
  fmt.Println("Visit Garden")
  // gimme := restorePlant(0)
  // gimme.LoadGens()
  // gimme.Grow(7)
  // gimme.Println(flower.Gens)
  // gimme.Println(flower.Render())
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

  fmt.Println(Plants)

  for true {

    prompt := ui.Select{
  		Label: "Garden Main Menu",
  		Items: []string{"Visit Garden", "Check Plant Dex", "View Gardener License", "Exit"},
    }
    i, _, _ := prompt.Run()
    switch i {
    case 0:
      visitGarden()
    case 1:
      checkPlantDex()
    case 2:
      viewGardenerLicense()
    case 3:
      os.Exit(0)
    default:
      fmt.Println("")
    }

  }

}
