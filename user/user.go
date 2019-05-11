package user

// +--------------+
// | Dependencies |
// +--------------+

import (
  "os"
  "strings"
  "io/ioutil"
)

// +---------+
// | Helpers |
// +---------+

// +-------------+
// | User struct |
// +-------------+

type User struct {
  SavePath string
  Plants [][]string
}

func (u User) GetSaveDir (string) {
  possibleConfigDirs := []string{"%APPDATA%", "${XDG_CONFIG_HOME}", "${HOME}/Library/Application"}
  for _, dir := range possibleConfigDirs {
    gimmeDir := string(dir)
    if _, err := os.Stat(dir); err == nil {
      return gimmeDir + "/garden/user/"
    }
  }
  return "no save path could be generated"
}

func (u *User) Load() {
  gimmePath := p.GetSaveDir() + "plants.sav"
  f, err := ioutil.ReadFile(gimmePath)
  if err != nil {
    u.Plants = [][]string{[]string{"No Plants Found"}}
  } else {
    gimmePlants := strings.Split(string(f), "\n")
    for i, gimmePlant := range gimmePlants {
      species, discriminator := strings.Split(string(gimmePlant), "  ")
      u.Plants[i] = []string{species, discriminator}
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
  for _, (species, discriminator) := range u.Plants {
    data = append(data, species + "  " + discriminator)
  }
  serializedData := strings.Join(data, "\n")
  message := []byte(serializedData)
  savePath := u.GetSaveDir() + "plants.sav"
  err := ioutil.WriteFile("testdata/hello", message, 0644)
}
