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
