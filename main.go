package main

// +--------------+
// | Dependencies |
// +--------------+

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "github.com/EthanThatOneKid/garden/static" // GetOsSaveDir, Plants ([]string)
  . "github.com/EthanThatOneKid/garden/plant" // Plant (struct)
  . "github.com/EthanThatOneKid/garden/user" // User (struct)
)

// +---------+
// | Globals |
// +---------+

var reader *bufio.Reader

// +---------+
// | Helpers |
// +---------+

func interpretSelection(options []string, input string) int {
  for i, option := range options {
    result := strings.Index(strings.ToUpper(option), input)
    if result > -1 {
      return i
    }
  }
  return -1
}

func restorePlant(species string) (Plant) {
  p := static.Plants[species]
  return Plant{
    Species: string(p[0]),
    Entry:   string(p[1]),
    Axiom:   string(p[2]),
    GrowthConfigX: strings.Split(string(p[3]), ","),
    GrowthConfigY: strings.Split(string(p[4]), ","),
    RenderConfigX: strings.Split(string(p[5]), ","),
    RenderConfigY: strings.Split(string(p[6]), ","),
  }
}

func getPlantSummary(data []string) (string) {
  species := string(data[0])
  p := restorePlant(species)
  p.SetDiscriminator(string(data[1]))
  p.LoadGens()
  return strconv.Itoa(p.GetGrowthLevel())
}

func handleMainMenuInput(selection int) {
  switch selection {
  case 0:
    visitGarden()
  case 1:
    checkPlantDex()
  case 2:
    viewGardenerLicense()
  case 3:
    os.Exit(0)
  default:
    fmt.Println("Try Again..?")
  }
}

func visitGarden() {
  u := User{}
  u.Load()
  var plantSelection []string
  for i, data := range u.Plants {
    species := string(data[0])
    discriminator := string(data[1])
    optionIndex := strconv.Itoa(i + 1)
    growthLevel := getPlantSummary(data)
    option := ("[" + optionIndex + "] " + species + " (" + discriminator + "), lv: " + growthLevel)
    plantSelection = append(plantSelection, option)
  }
  plantSelectionRender := "Choose a Plant:\n" + strings.Join(plantSelection, "\n")
  fmt.Println(plantSelectionRender)
  fmt.Printf("> ")
  byteName, _, _ := reader.ReadLine()
  inputValue := string(byteName)
  for i, data := range u.Plants {
    optionIndex := strconv.Itoa(i + 1)
    if (optionIndex == inputValue) {
      gimmeSpecies := string(data[0])
      gimmePlant := restorePlant(gimmeSpecies)
      gimmePlant.SetDiscriminator(string(data[1]))
      gimmePlant.LoadGens()
      viewPlant(gimmePlant)
    }
  }
}

func viewPlant(p Plant) {
  p.Grow(3)
  fmt.Println(p.Render())
}

func checkPlantDex() {
  fmt.Println("?Check plant dex")
}

func viewGardenerLicense() {
  fmt.Println("?viewGardenerLicense")
}

// +--------------+
// | Main Process |
// +--------------+

func main() {

  fmt.Println(static.Splash)
  reader = bufio.NewReader(os.Stdin)
  mainMenuOptions := []string{"Visit Garden", "Check Plant Dex", "View Gardener License", "Exit"}
  mainMenuOptionsRender := "Menu: [" + strings.Join(mainMenuOptions, "], [") + "]"

  for true {

    fmt.Println(mainMenuOptionsRender)
    fmt.Printf("> ")

    byteName, _, _ := reader.ReadLine()
    userInput := strings.ToUpper(string(byteName))
    fmt.Println()

    selection := interpretSelection(mainMenuOptions, userInput)
    handleMainMenuInput(selection)

    fmt.Println("\n")

  }

}
