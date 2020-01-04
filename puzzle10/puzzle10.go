package main

import (
    "fmt"
    "math"
    "os"
    "bufio"
    "sort"
)

type Asteroid struct {
  X int
  Y int
}

type AsteroidTarget struct {
  X int
  Y int
  Theta float64
  Distance float64
}

func main() {
  file, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  asteroid_map := []Asteroid{}
  y_coord := 0 // line number of file
  for scanner.Scan() {
    // By default, each token is separated by newline
    line := scanner.Text()

    // Construct asteroids and asteroid map (iterate over line)
    for x := 0; x < len(line); x++ {
      if line[x] == '#' {
        new_asteroid := Asteroid{x,y_coord}
        asteroid_map = append(asteroid_map, new_asteroid)
      }
    }
    y_coord++
  }

  // Now that we have our asteroid map, Let's find the best asteroid!
  // Calculate angle toward each other asteroid in asteroid map
  // Count unique angles, asteroid with max of this count is the best asteroid
  laser_base := Asteroid{0,0}
  max_angle_count := 0
  for _, asteroid1 := range asteroid_map{
    angle_set := make(map[float64]bool)
    for _, asteroid2 := range asteroid_map {
      x_diff := float64(asteroid2.X - asteroid1.X)
      y_diff := float64(asteroid2.Y - asteroid1.Y)
      angle := math.Atan2(y_diff, x_diff) * (180/math.Pi);
      angle_set[angle] = true
    }
    angle_count := len(angle_set)
    if angle_count > max_angle_count {
      max_angle_count = angle_count
      laser_base = asteroid1
    }
  }
  // Once we have our laser base, we need to find the 200th asteroid it blows up
  // This is easy. Compile a map of angle -> closest asteroid
  // Then sort it by angle 90 degrees being the zero. (if I understand polar coords right)
  asteroid_target_set := make(map[float64]AsteroidTarget)
  for _, asteroid2 := range asteroid_map {
    x_diff := float64(asteroid2.X - laser_base.X)
    y_diff := float64(asteroid2.Y - laser_base.Y)
    atan2 := math.Atan2(y_diff, x_diff) //+ (0.5*math.Pi)
    angle := atan2 * (180/math.Pi);
    if angle < 0 {
      angle = (atan2 + (2 * math.Pi)) * (180/math.Pi)
    }
    angle += 90.0
    if angle >= 360 {
      angle -= 360
    }
    // The initial firing angle is now 270 moving clockwise, so counting down
    distance_from_laser := math.Sqrt(x_diff*x_diff + y_diff*y_diff)
    asteroid_target := AsteroidTarget{asteroid2.X, asteroid2.Y, angle, distance_from_laser}
    if current_target, ok := asteroid_target_set[angle]; ok { // ISNT GOLANG COOL???
      if asteroid_target.Distance < current_target.Distance { // we have a closer asteroid
        asteroid_target_set[angle] = asteroid_target
      }
    } else { // its empty, put an asteroid to blow up in there
      asteroid_target_set[angle] = asteroid_target
    }
  }

  fmt.Println(asteroid_target_set[180])

  // We have constructed a set of >200 close asteroid targets.
  // We can either sort them or just blow them up while sorting them.
  // This is a lazy way (performance wise), blow up minimum every time

  // prev_target_angle := 360.0
  // for i := 0; i < 315; i++ {
  //   min_angle := -0.01
  //   current_target := AsteroidTarget{0,0,0.0,0.0}
  //   for angle, asteroid_target := range asteroid_target_set {
  //     if angle > min_angle && angle < prev_target_angle  {
  //       min_angle = angle
  //       current_target = asteroid_target
  //     }
  //   }
  //   fmt.Printf("%dth Asteroid: %v\n", i, current_target)
  //   prev_target_angle = min_angle
  // }
  key_slice := []float64{}
  for k, _ := range asteroid_target_set {
    key_slice = append(key_slice, k)
  }
  sort.Float64s(key_slice)
  fmt.Println(key_slice)
  fmt.Println(key_slice[199])
  fmt.Println(asteroid_target_set[key_slice[199]])
}
