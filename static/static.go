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
   (by EthanThatOneKid, devoted to GDDmad)`

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

    "Flowey": []string{
      "Flowey",
      "A test a test of a test yay",
      "a",
      "a,b,c",
      "bac,sb,cs",
      "a,b,c,s",
      "|,\\,/, ",
    },

    "Pine": {
      "Pine",
      "A pine tree",
      "a",
      "a,b,d,e",
      "bcdedcb,cc,ee,",
      "a,b,c,d,e",
      "||,|,|,|,|",
    },
}
