package plant

import (
  // "fmt"
  // "math"
  // "strings"
)

type Plant struct {
  Name string
  Gens, GrowthConfigX, GrowthConfigY []string
}

func (p *Plant) Phase(phaseConfigX []string, phaseConfigY []string) {
  for i := range p.Gens {
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
