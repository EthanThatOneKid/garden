package plant

// +--------------+
// | Dependencies |
// +--------------+

import (
  // "fmt"
  // "math"
  "os"
  "strings"
  "io/ioutil"
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
      extract = string(prevGen[0:len(from)])
      if extract == from {
        nextGen += to
        break
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

// +--------------+
// | Plant struct |
// +--------------+

type Plant struct {
  Species, Axiom string
  Discriminator int
  GrowthConfigX, GrowthConfigY []string
  RenderConfigX, RenderConfigY []string
  // PhaseConfigX,  PhaseConfigY  []string
  Gens []string
}

func (p Plant) GetSavePath() (string) {
  possibleConfigDirs := []string{"%APPDATA%", "${XDG_CONFIG_HOME}", "${HOME}/Library/Application"}
  for _, dir := range possibleConfigDirs {
    gimmeDir := string(dir)
    if _, err := os.Stat(dir); err == nil {
      return gimmeDir + "/" + p.Species + "/" + string(p.Discriminator) + "/"
    }
  }
  return "no save path could be generated"
}

func (p *Plant) LoadGens() {
  gimmePath := p.GetSavePath()
  f, err := ioutil.ReadFile(gimmePath)
  if err != nil {
    p.Gens = []string{p.Axiom}
  } else {
    p.Gens = strings.Split(string(f), "\n")
  }
}

func (p Plant) SaveGens() {
  gimmePath := p.GetSavePath()
  rawData := strings.Join(p.Gens, "\n")
  message := []byte(rawData)
  err := ioutil.WriteFile(gimmePath, message, 0644)
  if err != nil {
    os.Exit(0)
	}
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

func (p *Plant) Chop(layers int) {
  gimmeChoppingIndex := len(p.Gens) - layers - 1
  p.Gens = p.Gens[:gimmeChoppingIndex]
}
