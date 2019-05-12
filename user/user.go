package user

// +--------------+
// | Dependencies |
// +--------------+

import (
  //"os"
  "fmt"
  "strings"
  "io/ioutil"
  "../static"
)

// +---------+
// | Helpers |
// +---------+

// +-------------+
// | User struct |
// +-------------+

type User struct {
  Plants [][]string
}

func (u User) GetSaveDir() (string) {
  return static.GetOsSaveDir("/GARDEN/user/")
}

func (u *User) Load() {
  gimmePath := u.GetSaveDir() + "\\plants.sav"
  f, err := ioutil.ReadFile(gimmePath)
  if err != nil {
    u.Plants = [][]string{[]string{"No Plants Found"}}
  } else {
    gimmePlants := strings.Split(string(f), "\n")
    for _, gimmePlant := range gimmePlants {
      gimmePlantData := strings.Split(string(gimmePlant), "  ")
      species := string(gimmePlantData[0])
      discriminator := string(gimmePlantData[1])
      u.Plants = append(u.Plants, []string{species, discriminator})
    }
  }
}

func (u *User) Update(species, discriminator string) {
  gimmeNewPlant := []string{species, discriminator}
  u.Plants = append(u.Plants, gimmeNewPlant)
  u.Save()
}

func (u User) Save() {
  data := []string{}
  for _, plant := range u.Plants {
    species := string(plant[0])
    discriminator := string(plant[1])
    data = append(data, species + "  " + discriminator)
  }
  serializedData := strings.Join(data, "\n")
  message := []byte(serializedData)
  savePath := u.GetSaveDir() + "/plants.sav"
  err := ioutil.WriteFile(savePath, message, 0644)
  if err != nil {
    fmt.Println(err)
  }
}
