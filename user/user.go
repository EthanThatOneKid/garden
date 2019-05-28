package user

// +--------------+
// | Dependencies |
// +--------------+

import (
  //"os"
  //"fmt"
  "strings"
  "io/ioutil"
  "github.com/EthanThatOneKid/garden/static"
)

// +---------+
// | Helpers |
// +---------+

// +-------------+
// | User struct |
// +-------------+

type User struct {
  Plants [][]string
  PlantsSeen []string
}

func (u User) GetSaveDir() (string) {
  return static.GetOsSaveDir("/GARDEN/user/")
}

func (u *User) Load() {
  // Loading u.Plants
  gimmePath := u.GetSaveDir() + "\\plants.sav"
  f, err := ioutil.ReadFile(gimmePath)
  if err != nil {
    u.Plants = [][]string{}
  } else {
    gimmePlants := strings.Split(string(f), "\n")
    for _, gimmePlant := range gimmePlants {
      gimmePlantData := strings.Split(string(gimmePlant), "  ")
      species := string(gimmePlantData[0])
      discriminator := string(gimmePlantData[1])
      u.Plants = append(u.Plants, []string{species, discriminator})
    }
  }
  // Loading u.PlantsSeen
  gimmePath = u.GetOsSaveDir() + "\\seen.sav"
  f, err := ioutil.ReadFile(gimmePath)
  if err != nil {
    u.PlantsSeen = []string{}
  } else {
    u.PlantsSeen = strings.Split(string(f), "\n")
  }
}

func (u User) Has(species string) bool {
  for _, data := range u.Plants {
    candidate := string(data[0])
    if species == candidate {
      return true
    }
  }
  return false
}

func (u *User) RemovePlant(i int) {
  u.Plants = append(u.Plants[:i], u.Plants[i + 1:]...)
}

func (u *User) Update(species, discriminator string) {
  gimmeNewPlant := []string{species, discriminator}
  u.Plants = append(u.Plants, gimmeNewPlant)
  neverSeenSpecies := true
  for _, seenSpecies := range u.PlantsSeen {
    if species == seenSpecies {
      neverSeenSpecies = false
    }
  }
  if neverSeenSpecies {
    u.PlantsSeen = append(u.PlantsSeen, species)
  }
  u.Save()
}

func (u User) Save() {
  // Saving User.Plants
  data := []string{}
  for _, plant := range u.Plants {
    species := string(plant[0])
    discriminator := string(plant[1])
    data = append(data, species + "  " + discriminator)
  }
  serializedData := strings.Join(data, "\n")
  message := []byte(serializedData)
  savePath := u.GetSaveDir() + "/plants.sav"
  ioutil.WriteFile(savePath, message, 0644)
  // Saving User.PlantsSeen
  message = []byte(strings.Join(u.PlantsSeen, "\n"))
  savePath := u.GetSaveDir)_ + "/seen.sav"
  ioutil.WriteFile(savePath, message, 0644)

}
