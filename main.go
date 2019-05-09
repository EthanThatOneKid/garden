package main

// +--------------+
// | Dependencies |
// +--------------+

import (
  "os"
  "fmt"
  "sync"
  "strings"
  ui "github.com/manifoldco/promptui"
  // ui "gopkg.in/AlecAivazis/survey.v1"
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

func handleMainMenuInput(selection int, wg *sync.WaitGroup) {
  switch selection {
  case 0:
    go visitGarden(&wg)
  case 1:
    go checkPlantDex(&wg)
  case 2:
    go viewGardenerLicense(&wg)
  case 3:
    os.Exit(0)
  }
}

func visitGarden(wg *sync.WaitGroup) {
  fmt.Println("?Visit Garden")
  wg.Done()
}

func checkPlantDex(wg *sync.WaitGroup) {
  fmt.Println("?Check plant dex")
  wg.Done()
}

func viewGardenerLicense(wg *sync.WaitGroup) {
  fmt.Println("?viewGardenerLicense")
  wg.Done()
}

// +--------------+
// | Main Process |
// +--------------+

func main() {

  var wg sync.WaitGroup
  var selection int

  prompt := ui.Select{
    Label: "Garden Main Menu",
    Items: []string{"Visit Garden", "Check Plant Dex", "View Gardener License", "Exit"},
  }
  selection, _, _ = prompt.Run()

  wg.Add(1)
  go handleMainMenuInput(selection, &wg)
  wg.Wait()

  fmt.Println("again")

  //main()

}
