package plant

// +--------------+
// | Dependencies |
// +--------------+

import (
  "os"
  "strings"
  "io/ioutil"
  "math/rand"
  "time"
  "github.com/EthanThatOneKid/garden/static"
)

// +---------+
// | Helpers |
// +---------+

func generate(prevGen string, rulesX, rulesY []string) (string) {
  nextGen := ""
  for len(prevGen) > 0 {
    extract := ""
    for i, _ := range rulesX {
      from := string(rulesX[i])
      to := string(rulesY[i])
      if len(from) <= len(prevGen) {
        extract = string(prevGen[0:len(from)])
        if extract == from {
          nextGen += to
          break
        }
      }
      extract = ""
    }
    prevGen = string(prevGen[len(extract):])
    if len(extract) == 0 {
      nextGen += string(prevGen[:1])
      prevGen = string(prevGen[1:])
    }
  }
  return nextGen
}

func longestString(gens []string) (longest int) {
  longest = 0
  for _, v := range gens {
    genLen := len(string(v))
    if genLen > longest {
      longest = genLen
    }
  }
  return
}

func rndString(size int) (string) {
  alpha := "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
  rand.Seed(time.Now().UnixNano())
  buf := make([]byte, size)
  for i := 0; i < size; i++ {
    buf[i] = alpha[rand.Intn(len(alpha))]
  }
  return string(buf)
}

// +--------------+
// | Plant struct |
// +--------------+

type Plant struct {
  Species, Axiom, Entry, Discriminator string
  GrowthConfigX, GrowthConfigY []string
  RenderConfigX, RenderConfigY []string
  // PhaseConfigX,  PhaseConfigY  []string
  Gens []string
}

func (p *Plant) SetDiscriminator(disc string) {
  p.Discriminator = disc
}

func (p Plant) GetSaveDir() (string) {
  additionalDirs := "/GARDEN/plants/" + p.Species + "/" + string(p.Discriminator) + "/"
  return static.GetOsSaveDir(additionalDirs)
}

func (p *Plant) LoadGens() {
  gimmePath := p.GetSaveDir() + "/gens.sav"
  f, err := ioutil.ReadFile(gimmePath)
  if err != nil {
    p.Gens = []string{p.Axiom}
    if len(p.Discriminator) == 0 {
      p.Discriminator = rndString(5)
    }
  } else {
    p.Gens = strings.Split(string(f), "\n")
  }
}

func (p Plant) SaveGens() {
  gimmePath := p.GetSaveDir() + "/gens.sav"
  rawData := strings.Join(p.Gens, "\n")
  message := []byte(rawData)
  err := ioutil.WriteFile(gimmePath, message, 0644)
  if err != nil {
    os.Exit(0)
	}
}

func (p Plant) GetGrowthLevel() (int) {
  return len(p.Gens)
}

func (p Plant) Render() (string) {
  result := []string{}
  for i := len(p.Gens) - 1; i >= 0; i-- {
    gimmeRow := generate(p.Gens[i], p.RenderConfigX, p.RenderConfigY)
    result = append(result, gimmeRow)
  }
  halfLongestGen := longestString(result) / 2
  render := ""
  for i := 0; i < len(result); i++ {
    curGen := string(result[i])
    curLen := len(curGen)
    offset := int(halfLongestGen - (curLen / 2))
    render += strings.Repeat(" ", offset) + curGen + "\n"
  }
  return render
}

// func (p *Plant) Phase(layers int) {
//   from := len(p.Gens) - 1
//   to := from - layers
//   for i := from; i > to; i-- {
//     p.Gens[i] = generate(string(p.Gens[i]), p.PhaseConfigX, p.PhaseConfigY)
//   }
// }

func (p *Plant) Grow(epochs int) {
  prev := string(p.Gens[len(p.Gens) - 1])
  next := generate(prev, p.GrowthConfigX, p.GrowthConfigY)
  p.Gens = append(p.Gens, next)
  if epochs > 1 {
    p.Grow(epochs - 1)
  }
}

func (p *Plant) Trim(layers int) {
  gimmeChoppingIndex := len(p.Gens) - layers
  if gimmeChoppingIndex > 0 {
    p.Gens = p.Gens[:gimmeChoppingIndex]
  }
}
