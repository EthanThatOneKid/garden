package static

// +--------------+
// | Dependencies |
// +--------------+

import (
  "os"
  "fmt"
  "runtime"
  "path/filepath"
)

// +---------+
// | Helpers |
// +---------+

func GetOsSaveDir(additionalDirs string) (string) {
  var result string
  switch runtime.GOOS {
  case "windows":
    result = os.Getenv("APPDATA")
  case "darwin":
    result = os.Getenv("XDG_CONFIG_HOME")
  case "linux":
    result = os.Getenv("HOME") + "/Library/Application/"
  default:
    result = ""
  }
  result = filepath.Join(result, additionalDirs)
  if _, err := os.Stat(result); os.IsNotExist(err) {
    err := os.MkdirAll(result, os.ModeDir)
    if err != nil {
      fmt.Println(err)
    }
  }
  return result
}


// +-----------+
// | Constants |
// +-----------+

var Splash string = `
  .-_'''-.      ____    .-------.     ______         .-''-.  ,---.   .--.
 '_( )_   \   .'  __ '. |  _ _   \   |    _ '''.   .'_ _   \ |    \  |  |
|(_ o _)|  ' /   '  \  \| ( ' )  |   | _ | ) _  \ / ( ' )   '|  ,  \ |  |
. (_,_)/___| |___|  /  ||(_ o _) /   |( ''_'  ) |. (_ o _)  ||  |\_ \|  |
|  |  .-----.   _.-'   || (_,_).' __ | . (_) '. ||  (_,_)___||  _( )_\  |
'  \  '-   .'.'   _    ||  |\ \  |  ||(_    ._) ''  \   .---.| (_ o _)  |
 \  '-''   | |  _( )_  ||  | \ ''   /|  (_.\.' /  \  '-'    /|  (_,_)\  |
  \        / \ (_ o _) /|  |  \    / |       .'    \       / |  |    |  |
   ''-...-'   '.(_,_).' ''-'   ''-'  '-----''       ''-..-'  '--'    '--'
                (by EthanThatOneKid, devoted to GDDmad)
`

var Plants = map[string][]string{

    "Algae": []string{
      "Algae",
      "A simple, nonflowering, and typically aquatic plant of a large group that includes the seaweeds and many single-celled forms. Algae contain chlorophyll but lack true stems, roots, leaves, and vascular tissue.",
      "a",
      "a,b",
      "ab,a",
      "a,b",
      "\\,/",
    },

    "Rose": []string{
      "Rose",
      "A rose is a woody perennial flowering plant of the genus Rosa, in the family Rosaceae, or the flower it bears. There are over three hundred species and thousands of cultivars.",
      "a",
      "a,b,c",
      "bac,sb,cs",
      "a,b,c,s",
      "|,\\,/, ",
    },

    "Pine": {
      "Pine",
      "Pine trees are evergreen, coniferous resinous trees (or, rarely, shrubs) growing 3–80 m (10–260 ft) tall, with the majority of species reaching 15–45 m (50–150 ft) tall.",
      "a",
      "a,b,d,e",
      "bcdedcb,cc,ee,",
      "a,b,c,d,e",
      "||,|,|,|,|",
    },

    "Vine": {
      "Vine",
      "A vine is any plant with a growth habit of trailing or climbing stems, lianas or runners.",
      "a",
      "a,b,c,d",
      "b,c,d,a",
      "a,b,c,d",
      "<| | , |v| , | |>, | | ",
    },

/*

Bush #1
axiom = Y
X -> X[-FFF][+FFF]FX
Y -> YFX[+Y][-Y]
angle = 25.7

Bush #2
axiom = F
F -> FF+[+F-F-F]-[-F+F+F]
angle = 22.5

Bush #3
axiom = F
F -> F[+FF][-FF]F[-F][+F]F
angle = 35

Bush #4
axiom = VZFFF
V -> [+++W][---W]YV
W -> +X[-W]Z
X -> -W[+X]Z
Y -> YZ
Z -> [-FFF][+FFF]F
angle = 20

Tree
axiom = FX
X -> >[-FX]+FX
angle = 40

*/

}
