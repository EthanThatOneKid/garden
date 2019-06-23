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

    "Common Bush": []string{
      "Common Bush",
      "",
      "a",
      "a,b",
      "ab,a",
      "a,b",
      "\\,/",
    },

    "Tulip": []string{
      "Tulip",
      "",
      "a",
      "a,b,c",
      "bac,sb,cs",
      "a,b,c,s",
      "|,(,), ",
    },

    "Pine": []string{
      "Pine",
      "Pine trees are evergreen, coniferous resinous trees (or, rarely, shrubs) growing 3–80 m (10–260 ft) tall, with the majority of species reaching 15–45 m (50–150 ft) tall.",
      "a",
      "a,b,c,d,e,f",
      "b,c,d,e,f,g",
      "a,b,c,d,e,f,g",
      "|  |,/________\\,/______\\,/____\\,/__\\,/\\,",
    },

    "Vine": []string{
      "Vine",
      "A vine is any plant with a growth habit of trailing or climbing stems, lianas or runners.",
      "a",
      "a,b,c,d",
      "b,c,d,a",
      "a,b,c,d",
      "<| | , |v| , | |>, | | ",
    },

    "Bamboo": []string{
      "Bamboo",
      "Bamboo has a specific tensile strength that rivals steel. In Philippine mythology, one of the more famous creation accounts tells of the first man and woman, Maganda, each emerged from one half of a split bamboo stem on an island formed after the battle between Sky and Ocean.",
      "a",
      "a,b",
      "b,a",
      "a,b",
      "|__|,|`.|",
    },

    "Succulent": []string{
      "Succulent",
      "Succulents are plants that have some parts that are more than normally thickened and fleshy, usually to retain water in arid climates or soil conditions.",
      "a",
      "a,b",
      "b,a",
      "a,b",
      "{#( , )#}",
    },

    "Cactus #1": []string{
      "Cactus #1",
      "",
      "a",
      "a,b,lsr",
      "b,lsr,asssa",
      "a,b,l,r,s",
      "|.|,\\ V /,\\:\\,/:/, ",
    },

    "Cactus #2": []string{
      "Cactus #2",
      "",
      "a",
      "a,b,c,d,e,f",
      "b,c,d,e,f,a",
      "a,b,c,d,e,f",
      "    | |    ,    |  ___),    | |_( ),(___  |  * ,( )_| |    , *  | |    ",
    },

    "Kelp": []string{
      "Kelp",
      "Kelp are large brown algae that live in cool, relatively shallow waters close to the shore. They grow in dense groupings much like a forest on land. These underwater towers of kelp provide food and shelter for thousands of fish, invertebrates, and marine mammal species.",
      "a",
      "a,b,c,d",
      "b,c,d,a",
      "a,b,c,d",
      "( : . (  , / . : / ,  ) : . ), \\ . : \\ ",
    },

    "Bonzai": []string{
      "Kelp",
      "Kelp are large brown algae that live in cool, relatively shallow waters close to the shore. They grow in dense groupings much like a forest on land. These underwater towers of kelp provide food and shelter for thousands of fish, invertebrates, and marine mammal species.",
      "a",
      "a,b,c,d,e,f,j,h,i",
      "b,c,d,e,f,g,h,i,e",
      "a,b,c,d,e,f,g,h,i",
      "          /)#(         ,         _.) ;/ *%;     ,     *%    \\/ #);-\"*%%* ,%^     ;*%%% )\\|.%%*%;_,  % *%%; ;%%%%*(    '   ,  ;%%%       \\(_.*%%%%. , ;%%%%%*%   _%%%%\"      ,       ;'%% \\-*%%%%%%%,    %%%;%%%%%%%         ",
    },

    // "Palm Tree": []string{
    //   "Kelp",
    //   "Kelp are large brown algae that live in cool, relatively shallow waters close to the shore. They grow in dense groupings much like a forest on land. These underwater towers of kelp provide food and shelter for thousands of fish, invertebrates, and marine mammal species.",
    //   "a",
    //   "a,b,c,d",
    //   "b,c,d,a",
    //   "a,b,c,d",
    //   "( : . (  , / . : / ,  ) : . ), \\ . : \\ ",
    // },
    //
    // "X-mas Tree": []string{
    //   "Kelp",
    //   "Kelp are large brown algae that live in cool, relatively shallow waters close to the shore. They grow in dense groupings much like a forest on land. These underwater towers of kelp provide food and shelter for thousands of fish, invertebrates, and marine mammal species.",
    //   "a",
    //   "a,b,c,d",
    //   "b,c,d,a",
    //   "a,b,c,d",
    //   "( : . (  , / . : / ,  ) : . ), \\ . : \\ ",
    // },
    //
    // "Toadstool": []string{
    //   "Kelp",
    //   "Kelp are large brown algae that live in cool, relatively shallow waters close to the shore. They grow in dense groupings much like a forest on land. These underwater towers of kelp provide food and shelter for thousands of fish, invertebrates, and marine mammal species.",
    //   "a",
    //   "a,b,c,d",
    //   "b,c,d,a",
    //   "a,b,c,d",
    //   "( : . (  , / . : / ,  ) : . ), \\ . : \\ ",
    // },
    //
    // "Common Weed": []string{
    //   "Kelp",
    //   "Kelp are large brown algae that live in cool, relatively shallow waters close to the shore. They grow in dense groupings much like a forest on land. These underwater towers of kelp provide food and shelter for thousands of fish, invertebrates, and marine mammal species.",
    //   "a",
    //   "a,b,c,d",
    //   "b,c,d,a",
    //   "a,b,c,d",
    //   "( : . (  , / . : / ,  ) : . ), \\ . : \\ ",
    // },




// Palm Tree
//      __ _.--..--._ _
//   .-' _/   _/\_   \_'-.
//  |__ /   _/\__/\_   \__|
//     |___/\_\__/  \___|
//            \__/
//            \__/
//             \__/
//              \__/
//           ____\__/___

// Christmas Tree
  //     _\/_
  //      /\
  //      /\
  //     /  \
  //     /~~\o
  //    /o   \
  //   /~~*~~~\
  //  o/    o \
  //  /~~~~~~~~\~`
  // /__*_______\
  //      ||

// Toadstool
//   .-"""-.
//  /* * * *\
// :_.-:`:-._;
//     (_)
//  \|/(_)\|/

// Weed
 //       |
 //     \|/|/
 //   \|\\|//|/
 //    \|\|/|/
 //     \\|//
 //      \|/
 //      \|/
 //       |
 // _\|/__|_

 // Cactus #3
 //      *-*,
 // ,*\/|`| \
 // \'  | |'| *,
 //  \ `| | |/ )
 //   | |'| , /
 //   |'| |, /
 // __|_|_|_|__

 // Giant Sequoia
//                      |_    /
//                      |-   |
//                      |   =|
//                      |    |
// --------------------/ ,  . \

}
