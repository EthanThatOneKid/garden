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

func (p Plant) Grow() ([]string) { // https://en.wikipedia.org/wiki/L-system
  prev := p.Generations[len(p.Generations) - 1]
  fmt.Println(prev)
  next := ""
  for i := range prev {
    cur := string(prev[i])
    rule_loop:
    for j := range p.GrowthInstructions {
      from := string(p.GrowthInstructions[j][0])
      to := string(p.GrowthInstructions[j][1])
      if cur == from {
        cur = to
        break rule_loop
      }
    }
    next += cur
  }
  return append(p.Generations, next)
}
