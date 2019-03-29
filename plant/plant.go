package plant

import (
  // "fmt"
  // "math"
  "strings"
)

func getOffset(prevLen, prevOff, curLen int) (genOff int) {
  genOff = prevOff + int((prevLen / 2) - (curLen / 2))
  if genOff < 0 {
    genOff = 0
  }
  return
}

type Plant struct {
  Name string
  Gens, GrowthConfigX, GrowthConfigY []string
}

func (p Plant) Print() (string) {
  result := ""
  prevGenLen, prevOff := 0, 0
  for i := len(p.Gens) - 1; i >= 0; i-- {
    gimmeRow := ""
    chars := strings.Split(string(p.Gens[i]), "")
    // TODO : for each character, see what translate to for rendering bc rn its hardcoded
    for j := range chars {
      gimme := string(chars[j])
      if gimme == "A" {
        gimmeRow += "\\"
      }
      if gimme == "B" {
        gimmeRow += "/"
      }
      if gimme == "F" {
        gimmeRow += "o"
      }
    }
    curLen := len(p.Gens[i])
    genOffset := getOffset(prevGenLen, prevOff, curLen)
    prevGenLen = curLen
    prevOff = genOffset
    result += strings.Repeat(" ", genOffset) + gimmeRow + "\n"
  }
  return result
}

func (p *Plant) Phase(phaseConfigX []string, phaseConfigY []string, layers int) {
  from := len(p.Gens) - 1
  to := from - layers
  for i := from; i > to; i-- {
    accumulator := ""
    for j := range p.Gens[i] {
      cur := string(p.Gens[i][j])
      rule_loop: for k := range phaseConfigX {
        from := string(phaseConfigX[k])
        to := string(phaseConfigY[k])
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
