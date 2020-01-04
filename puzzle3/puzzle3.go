package main

import (
    "fmt"
    "strings"
    "strconv"
    "math"
    // "bufio"
)
type Vertex struct {
  X int
  Y int
}

type Vector struct {
	X1 int
  Y1 int
  X2 int
	Y2 int
}

// Function to create a new vector given a path and old vector
func MakeNextVector(path string, prev Vector) Vector {
  x_change := 0
  y_change := 0
  if strings.Contains(path, "R") {
    x_change, _ = strconv.Atoi(strings.TrimLeft(path, "R"))
  } else if strings.Contains(path, "L") {
    x_change, _ = strconv.Atoi(strings.TrimLeft(path, "L"))
    x_change = x_change * -1
  } else if strings.Contains(path, "U") {
    y_change, _ = strconv.Atoi(strings.TrimLeft(path, "U"))
  } else if strings.Contains(path, "D") {
    y_change, _ = strconv.Atoi(strings.TrimLeft(path, "D"))
    y_change = y_change * -1
  }
  return Vector{prev.X2, prev.Y2, prev.X2 + x_change, prev.Y2 + y_change}
}

func HasIntersection(v1 Vector, v2 Vector) (bool, Vertex) {
  if (v1.X1 - v1.X2) == 0 { // vector 1 y value is changed
    if (v2.X1 - v2.X2) == 0 { // lines are parallel, no intersection
      return false, Vertex{0,0}  // vectors move in parallel
    } else { // perpendicular vectors, check for overlap
      var y_cross bool = ((v1.Y1 >= v2.Y1) && (v1.Y2 <= v2.Y1)) || ((v1.Y1 <= v2.Y1) && (v1.Y2 >= v2.Y1))
      var x_cross bool = ((v2.X1 >= v1.X1) && (v2.X2 <= v1.X1)) || ((v2.X1 <= v1.X1) && (v2.X2 >= v1.X1))
      if y_cross && x_cross  {
        intersection := Vertex{v1.X1, v2.Y1}
        return true, intersection // intersection detected
      }
    }
  } else if (v2.Y1 - v2.Y2) != 0 { // perpendicular vector
    var y_cross bool = ((v1.Y1 >= v2.Y1) && (v1.Y1 <= v2.Y2)) || ((v1.Y1 <= v2.Y1) && (v1.Y1 >= v2.Y2))
    var x_cross bool = ((v2.X1 >= v1.X1) && (v2.X1 <= v1.X2)) || ((v2.X1 <= v1.X1) && (v2.X1 >= v1.X2))
    if y_cross && x_cross {
      intersection := Vertex{v2.X1, v1.Y1}
      return true, intersection // intersection detected
    }
  }
  return false, Vertex{0,0} // no possible intersection
}

func GetDistanceFromOrigin(v Vertex) int {
  return int(math.Abs(float64(v.X))) + int(math.Abs(float64(v.Y)))
}

/*
Method added to calculate part 2
Being lazy and not optimizing here, assumes no overlapping on each wire
*/
func FindIntersectionDistance(intersection Vertex, w1_vectors []Vector, w2_vectors []Vector) int {
  w1_distance := 0
  w2_distance := 0
  for _, v := range w1_vectors {
    steps := int(math.Abs(float64(v.X2 - v.X1))) + int(math.Abs(float64(v.Y2 - v.Y1)))
    w1_distance += steps
  }
  for _, v := range w2_vectors {
    steps := int(math.Abs(float64(v.X2 - v.X1))) + int(math.Abs(float64(v.Y2 - v.Y1)))
    w2_distance += steps
  }
  last_v1 := w1_vectors[len(w1_vectors) - 1]
  last_v2 := w2_vectors[len(w2_vectors) - 1]
  w1_diff := int(math.Abs(float64(last_v1.Y2 - intersection.Y)) + math.Abs(float64(last_v1.X2 - intersection.X)))
  w2_diff := int(math.Abs(float64(last_v2.Y2 - intersection.Y)) + math.Abs(float64(last_v2.X2 - intersection.X)))
  w1_distance -= w1_diff
  w2_distance -= w2_diff
  return w1_distance + w2_distance
}

/*
Approach
Create vector slice for first wire, then check each new vector in second wire if
it crosses a vector from the first wire's list

If crossed, add intersection point to a list
Calculate running min in that list
*/
func main() {
  wire1_vectors := make([]Vector, 0)
  wire2_vectors := make([]Vector, 0)
  prev_vector := Vector{0,0,0,0}
  // wire1_pathing := []string{"R8","U5","L5","D3"}
  // wire2_pathing := []string{"U7","R6","D4","L4"}
  wire1_pathing := []string{"R995","D933","L284","U580","R453","U355","L352","U363","L506","D130","R300","D112","L751","U245","R174","U901","L586","D70","L348","U307","R596","D401","R311","D328","L612","D214","L161","U488","L813","U298","L299","D807","L791","D813","R946","U958","R258","D972","L72","U698","L170","D131","L552","D14","L760","U812","L203","D92","R483","U698","R800","U345","L947","D206","L991","D447","R832","D52","L134","D946","R108","D477","L818","D101","R784","U409","R23","U359","R981","D458","R786","U445","R874","U244","R381","U206","R86","U279","L787","U683","R52","U740","R210","U162","L748","U747","R283","D964","R782","D386","R205","D898","R774","U421","R908","D280","L567","D667","L302","D795","L822","D908","L160","U279","L58","D914","R828","U621","R338","U982","R471","U724","L509","U444","R377","D473","R307","U331","L529","U855","L474","U725","L905","U409","L881","U702","R162","D717","R498","U458","R119","U744","R2","D82","R337","D988","L235","U831","L373","D122","L310","D107","R685","U864","L356","U213","R246","U373","L809","U593","R582","U488","R847","U792","L182","U219","L232","D695","R183","U164","L730","D660","L581","D931","R36","D291","R614","U408","R928","D93","L685","D879","R37","D113","L563","D340","L212","D907","L557","D522","L610","D927","R11","U556","R571","U668","L834","U603","L407","U966","R245","D408","R822","U163","L265","D230","L144","D786","R479","U208","L781","D787","L146","U476","R561","U70","R879","U562","R805","D897","L44","U709","L773","U747","L806","U140","R953","D628","L752","D666","R916","D351","R175","D504","R232","U463","R415","U492","L252","D30","L574","U100","L333","U313","R403","D68","L976","D702","L205","D992","L552","U55","R216","U548","L894","U764","L919","U475","R656","U712","L754","U638","R33","U709","R643","U299","R151","U762","R767","D542","R581","D752","L701","D589","L312","U189","R922","D915","R975","U26","R244","U48","L971","U755","R889","D726","R126","U978","L275","D435","L732","D982","L654","D704","L98","U153","L983","U770","L429","U530","L545","D44","L36","U456","R442","D58","L321","U473","R657","U307","R314"}
  wire2_pathing := []string{"L999","U880","L754","D995","R105","U651","R762","U451","R612","U424","L216","D210","L946","U57","R685","U185","R234","D768","L927","U592","R545","U770","R456","D118","R22","U371","L721","D937","R144","U173","R337","D17","L948","U720","R617","D588","L57","U258","L979","U232","L473","D451","L829","D85","L985","D333","L492","D749","L972","U188","R214","D108","R247","U379","L218","D490","R451","U582","R674","D881","R725","U404","L916","U137","R745","D969","L206","D480","R634","U672","R897","D184","L768","D766","L529","U317","L909","D74","R529","D31","R483","D906","R961","D535","L937","D751","L564","U478","L439","U556","R385","D590","L518","D991","L232","D358","L962","U242","R856","U66","L847","U146","R196","U515","L892","U18","L535","U343","R430","U77","L967","U956","L64","D510","L29","D305","L954","U186","R624","D693","R354","D243","L145","D622","R456","U904","L233","D940","L82","D83","L726","D993","L338","D793","R340","D823","R147","D595","R296","D388","L106","D362","R114","U808","L496","U614","L809","U911","R480","D29","L802","U900","R920","U952","R237","D383","L480","U362","L761","U949","L920","D82","L511","U365","R657","U465","L256","U124","L810","U43","L668","U146","L847","D245","R937","D778","L995","D316","R423","U515","L626","D788","R453","U98","R886","U821","R749","D365","R915","U980","R322","D114","L556","D921","L551","D397","R232","D485","L842","D455","R940","U913","L196","D239","L221","D200","R438","U71","L979","U527","L86","U959","R768","D557","R553","D709","L838","U191","L916","D703","L687","U34","R463","D809","L3","U335","L231","D212","L674","U177","L51","D557","L939","U390","L28","U53","L653","D182","R329","D449","L563","D476","R258","D299","L142","U847","L38","U322","R294","U320","R314","D257","R622","U59","R501","D677","L880","U589","R599","D245","L851","U369","R262","D63","R722","D253","L546","U578","R909","U678","L63","U256","L237","U798","R465","D420","R797","D452","R548","U875","R523","D527","L467","U49","R726","D840","R882","U97","L398","D621","R38","U952","R196","D597","R627","D721","L441","D710","L18","D679","R218"}
  intersections := make([]Vertex, 0)
  for _, path := range wire1_pathing {
    next_vector := MakeNextVector(path, prev_vector)
    wire1_vectors = append(wire1_vectors, next_vector)
    prev_vector = next_vector
  }
  prev_vector = Vector{0,0,0,0}
  for _, path := range wire2_pathing {
    next_vector := MakeNextVector(path, prev_vector)
    wire2_vectors = append(wire2_vectors, next_vector)
    prev_vector = next_vector
  }
  min_time := 1000000000 // Go doesn't have a max_int, bit shifting is dumb. It's 2019
  fmt.Println(min_time)
  for i, wire2_vector := range wire2_vectors {
    // Check if new vector cross any vector in wire 1's vectors
    for j, wire1_vector := range wire1_vectors {
      isIntersection, intersection := HasIntersection(wire1_vector, wire2_vector)
      if isIntersection {
        intersections = append(intersections, intersection)
        distance := FindIntersectionDistance(intersection, wire1_vectors[0:j+1], wire2_vectors[0:i+1])
        if (distance < min_time) && (distance > 0) {
          min_time = distance
        }
      }
    }
  }
  fmt.Println(min_time)
  min := GetDistanceFromOrigin(intersections[1])
  fmt.Println(intersections)
  for _, intersection := range intersections {
    distance := GetDistanceFromOrigin(intersection)
    if (distance < min) && (distance != 0){
      min = distance
    }
  }

  fmt.Println(min)
}
