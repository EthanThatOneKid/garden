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
      "Shrubs are perennial woody plants, and therefore have persistent woody stems above ground. Usually shrubs are distinguished from trees by their height and multiple stems.",
      "a",
      "a,b",
      "ab,a",
      "a,b",
      "\\,/",
    },

    "Tulip": []string{
      "Tulip",
      "Tulips originally were found in a band stretching from Southern Europe to Central Asia, but since the seventeenth century have become widely naturalised and cultivated. In their natural state they are adapted to steppes and mountainous areas with temperate climates. Flowering in the spring, they become dormant in the summer once the flowers and leaves die back, emerging above ground as a shoot from the underground bulb in early spring.",
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
      "asbsasb",
      "a,b",
      "b,a",
      "a,b,s",
      "{#( , )#}, ",
    },

    "Cactus #1": []string{
      "Cactus #1",
      "Cacti have spines instead of leaves. Spines can be soft or rigid, straight or curved, arranged in rows or scattered. They can reach 6 inches in length.",
      "a",
      "a,b,lsr",
      "b,lsr,asssa",
      "a,b,l,r,s",
      "|.|,\\ V /,\\:\\,/:/, ",
    },

    "Cactus #2": []string{
      "Cactus #2",
      "Since cacti live in dry areas, they need to absorb a large amount of water and store it in the stem and roots for the periods of drought. Besides storing of water, the stem plays a significant role in the process of photosynthesis.",
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

    "Bonsai": []string{
      "Bonsai",
      "'Bonsai' refers to miniaturized, container-grown trees adhering to Japanese tradition and principles. The Japanese loanword 'bonsai' has become an umbrella term in English, attached to many unrelated forms of potted or other plants.",
      "a",
      "a,b,c,d,e,f,g,h,i",
      "b,c,d,e,f,g,h,i,e",
      "a,b,c,d,e,f,g,h,i",
      "/)#(,    .) ;/ *%;,   *%    \\/ #).-\"*%%*,%^    ;*%%% )\\|;%%*%;_,%*%%; ;%%%%*(    ',;%%%       \\(_.*%%%%.,;%%%%%*%   _%%%%\"     ,;'%% \\\\-*%%%%%%%,%%%;%%%%%%%",
    },

    "Palm Tree": []string{
      "Palm Tree",
      "Most palms are distinguished by their large, compound, evergreen leaves, known as fronds, arranged at the top of an unbranched stem. However, palms exhibit an enormous diversity in physical characteristics and inhabit nearly every type of habitat within their range, from rainforests to deserts.",
      "a",
      "a,b,c,d,e,f,g",
      "b,c,d,e,f,g,e",
      "a,b,c,d,e,f,g,h",
      "\\__/,\\__/  ,\\__/    ,\\__/     ,|___/\\_\\__/  \\___|,|__ /   _/\\__/\\_   \\__|,.-' _/   _/\\_   \\_'-.,__ _.--..--._ _",
    },

    "Daisy": []string{
      "Daisy",
      "In fellow Asteraceaean plants, what appears to be a single flower is actually a cluster of much smaller flowers. The overall appearance of the cluster, as a single flower, functions in attracting pollinators in the same way as the structure of an individual flower in some other plant families.",
      "a",
      "a,b,c,d,e,f,g,h,i,j,k,l,m,n",
      "b,c,d,e,f,g,h,i,j,k,l,m,n,a",
      "a,b,c,d,e,f,g,h,i,j,k,l,m,n",
      "              |/__/   '-'     ,              Y ___'-'   '-'  ,            \\|;' (   )-(   ),           \\ |  ..>-(   )-< ,           / \\   (   )-(   ),             .|    .-.   .-.  ,             |       .-.     ,  '-'   \\__\\|             , '-'   '-'___ Y               ,(   )-(   ) ';|/              , >-(   )-<..  | /             ,(   )-(   )   / \\             , .-.   .-.    |.              ,    .-.       |               ",
    },

    "Toadstool": []string{
      "Toadstool",
      "A mushroom develops from a nodule, or pinhead, less than two millimeters in diameter, called a primordium, which is typically found on or near the surface of the substrate. It is formed within the mycelium, the mass of threadlike hyphae that make up the fungus.",
      "a",
      "a,b,c,d,e,s",
      "b,c,d,e,s,s",
      "a,b,c,d,e,s",
      "\\|/(_)\\|/,(_),:_.-:':-._;,/* * * *\\,.-\"\"\"-.,",
    },

    "Giant Sequoia": []string{
      "Giant Sequoia",
      "Because of its size, the tree has been studied for its water pull. Water from the roots can be pushed up only a few meters by osmotic pressure but can reach extreme heights by using a system of branching capillarity in the tree's xylem and sub-pressure from evaporating water at the leaves.",
      "a",
      "a,b,c,d,e,f,g,h,i",
      "b,c,d,e,f,g,h,i,b",
      "a,b,c,d,e,f,g,h,i",
      "/ / 7 .  \"# \\, _/   :;  |,| o     ==|,|-  (<')) |,|_        /,|   __   /_,| (('>)   |,|--    \"\" |,|   6     |",
    },

    "Cactus #3": []string{
      "Cactus #3",
      "Certain cacti produce a substance, called mescaline, which induces hallucinogenic effects. It has been used by shamans to induce trans-like state and ensure 'communication' with G-d.",
      "a",
      "a,b,c,d,e",
      "b,c,d,e,a",
      "a,b,c,d,e",
      "  | |'| ; / , \\ :| | |/ ),\\'  | |:| *;,;*\\/|:| \\   ,  \\'|*-*;/  ",
    },

}
