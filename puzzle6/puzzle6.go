package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
)

var orbit_map = make(map[string]string)

func OrbitMapWalk(object string) int {
  center_of_mass, ok := orbit_map[object]
  if !ok { // No further orbits
    return 0
  } else {
    return 1 + OrbitMapWalk(center_of_mass)
  }
}

func OrbitMapWalk2(object string) []string {
  center_of_mass, ok := orbit_map[object]
  if !ok { // No further orbits
    return []string{center_of_mass}
  } else {
    return append(OrbitMapWalk2(center_of_mass), object)
  }
}

func GetIndirectOrbits(objects []string) int {
  total_indirect := 0
  for _, object := range objects {
    total_indirect += OrbitMapWalk(object)
  }
  return total_indirect
}
// elem, ok = m[key]

func BuildSantaMap() []string {
  // fmt.Println("here?")
  return OrbitMapWalk2("SAN")
}

func BuildMeMap() []string {
  // fmt.Println("here?")
  return OrbitMapWalk2("YOU")
}

func FindOrbitalTransfers(santa_map []string, me_map []string) int {
  for i := 0; (i < len(santa_map)) &&  (i < len(me_map)); i++ {
    if santa_map[i] != me_map[i] {
      santa_count := len(santa_map) - i - 1
      me_count := len(me_map) - i - 1
      fmt.Println(santa_count)
      fmt.Println(me_count)
      return me_count + santa_count
    }
  }
  return 0
}

/*
Build a map where all objects are keys and the object they directly orbit is the value
*/
func main() {
  // Open file with input data.
  file, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    // By default, each token is separated by newline
    orbit_val := scanner.Text()
    split_orbit_val := strings.SplitN(orbit_val, ")", 2)
    center_of_mass := split_orbit_val[0]
    orbiter := split_orbit_val[1]
    orbit_map[orbiter] = center_of_mass
  }
  // objects := make([]string, 0, len(orbit_map))
  // for k := range orbit_map {
  //   objects = append(objects, k)
  // }
  santa_map := BuildSantaMap()
  me_map := BuildMeMap()
  transfers := FindOrbitalTransfers(santa_map, me_map)
  fmt.Println(transfers)
  //
  // Solution for first puzzle. I messed up and somehow my method for
  // counting indirect orbits seemed to catch all of 'em
  // direct_orbits := 0
  // indirect_orbits := GetIndirectOrbits(objects)
  // fmt.Println(direct_orbits + indirect_orbits)

}
