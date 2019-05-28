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
  PlantsSeen map[string]bool
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
  u.PlantsSeen = make(map[string]bool)
  gimmePath = u.GetSaveDir() + "\\seen.sav"
  f, err = ioutil.ReadFile(gimmePath)
  if err == nil {
    gimmePlantsSeenList := strings.Split(string(f), "\n")
    for _, _species := range gimmePlantsSeenList {
      u.PlantsSeen[string(_species)] = true
    }
  }
  if err != nil || len(u.PlantsSeen) == 0 {
    for _, data := range u.Plants {
      u.PlantsSeen[string(data[0])] = true
    }
    u.Save()
  }
}

func (u User) Has(species string) bool {
  for _species, _ := range u.PlantsSeen {
    candidate := string(_species)
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
  u.PlantsSeen[species] = true
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
  var gimmePlantsSeenList []string
  for species, _ := range u.PlantsSeen {
    gimmePlantsSeenList = append(gimmePlantsSeenList, species)
  }
  message = []byte(strings.Join(gimmePlantsSeenList, "\n"))
  savePath = u.GetSaveDir() + "/seen.sav"
  ioutil.WriteFile(savePath, message, 0644)

}
