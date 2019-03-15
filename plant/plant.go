package plant

import (
  "fmt"
  // "math"
  // "strings"
)

type Plant struct {
  Name string
  Generations []string
  GrowthInstructions [][]string
}

func (p Plant) Grow() { // https://en.wikipedia.org/wiki/L-system
  prev := p.Generations[len(p.Generations) - 1]
  fmt.Println(prev)
  next := ""
  for i := range prev {
    cur := string(prev[i])
    //rule_loop:
    for j := range p.GrowthInstructions {
      from := p.GrowthInstructions[j][0]
      to := p.GrowthInstructions[j][1]
      if cur == from { // "mismatched types byte and string"
        cur = to
        // break rule_loop
      }
    }
    next += cur
  }
  p.Generations = append(p.Generations, next)
}
