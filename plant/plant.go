package plant

import (
  // "fmt"
  // "math"
  "strings"
  "io/ioutil"
)

func getOffset(prevLen, prevOff, curLen int) (genOff int) {
  genOff = prevOff + int((prevLen / 2) - (curLen / 2))
  if genOff < 0 {
    genOff = 0
  }
  return
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

// TODO : https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
type Plant struct {
  Species, Axiom string
  GrowthConfigX, GrowthConfigY []string
  RenderConfigX, RenderConfigY []string
  PhaseConfigX,  PhaseConfigY  []string
  Gens []string
}

func (p *Plant) LoadGens(path string) {
  f, err := ioutil.ReadFile(path)
  if err != nil {
    p.Gens = []string{p.Axiom}
  } else {
    p.Gens = strings.Split(string(f), " ")
  }
}

func (p Plant) Render() (string) {
  result := ""
  prevGenLen, prevOff := 0, 0
  for i := len(p.Gens) - 1; i >= 0; i-- {
    gimmeRow := ""
    chars := strings.Split(string(p.Gens[i]), "")
    for j := range chars {
      cur := string(chars[j])
      rule_loop: for k := range p.RenderConfigX {
        from := string(p.RenderConfigX[k])
        to := string(p.RenderConfigY[k])
        if cur == from {
          gimmeRow += to
          break rule_loop
        }
      }
    }
    curLen := len(gimmeRow)
    genOffset := getOffset(prevGenLen, prevOff, curLen)
    prevGenLen = curLen
    prevOff = genOffset
    result += strings.Repeat(" ", genOffset) + gimmeRow + "\n"
  }
  longestGen = longestString(res)
  //for i := len()
  return result
}

func (p *Plant) Phase(layers int) {
  from := len(p.Gens) - 1
  to := from - layers
  for i := from; i > to; i-- {
    accumulator := ""
    for j := range p.Gens[i] {
      cur := string(p.Gens[i][j])
      rule_loop: for k := range p.PhaseConfigX {
        from := string(p.PhaseConfigX[k])
        to := string(p.PhaseConfigY[k])
        if cur == from {
          accumulator += to
          break rule_loop
        } else {
          accumulator += cur
        }
      }
    }
    p.Gens[i] = accumulator
  }
}

func (p *Plant) Grow(epochs int) {
  prev := p.Gens[len(p.Gens) - 1]
  next := ""
  for i := range prev {
    cur := string(prev[i])
    rule_loop: for j := range p.GrowthConfigX {
      from := string(p.GrowthConfigX[j])
      to := string(p.GrowthConfigY[j])
      if cur == from {
        cur = to
        break rule_loop
      }
    }
    next += cur
  }
  p.Gens = append(p.Gens, next)
  if epochs > 1 {
    p.Grow(epochs - 1)
  }
}
