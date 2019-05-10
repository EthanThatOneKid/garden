package main

// +--------------+
// | Dependencies |
// +--------------+

import (
  "os"
  "fmt"
  //"sync"
  "bufio"
  "strings"
  //ui "github.com/manifoldco/promptui"
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

func interpretSelection(options []string, input string) int {
  for i, option := range options {
    result := strings.Index(strings.ToUpper(option), input)
    if result > -1 {
      return i
    }
  }
  return -1
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
    fmt.Println("Try Again...")
  }
}

func visitGarden() {
  fmt.Println("?Visit Garden")
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

  reader := bufio.NewReader(os.Stdin)
  mainMenuOptions := []string{"Visit Garden", "Check Plant Dex", "View Gardener License", "Exit"}

  for true {

    mainMenuOptionsRender := "? [" + strings.Join(mainMenuOptions, "], [") + "]"
    fmt.Println(mainMenuOptionsRender)
    fmt.Printf("> ")
    byteName, _, _ := reader.ReadLine()
    userInput := strings.ToUpper(string(byteName))
    selection := interpretSelection(mainMenuOptions, userInput)
    handleMainMenuInput(selection)
    fmt.Println("\n")

  }

}
