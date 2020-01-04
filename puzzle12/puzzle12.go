package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
)

type Moon struct {
  pos Position
  vel Velocity
}

type Position struct {
  px int
  py int
  pz int
}

type Velocity struct {
  vx int
  vy int
  vz int
}

func ApplyGravityOnAxis(a int, b int) (int, int) {
  if a < b {
    return 1, -1
  } else if a > b {
    return -1, 1
  } else {
    return 0, 0
  }
}

// Something about this function could be condensed.
func ApplyGravity(a Moon, b Moon) (Moon, Moon) {
  grav_ax, grav_bx := ApplyGravityOnAxis(a.pos.px, b.pos.px)
  grav_ay, grav_by := ApplyGravityOnAxis(a.pos.py, b.pos.py)
  grav_az, grav_bz := ApplyGravityOnAxis(a.pos.pz, b.pos.pz)
  a.vel.vx += grav_ax
  b.vel.vx += grav_bx
  a.vel.vy += grav_ay
  b.vel.vy += grav_by
  a.vel.vz += grav_az
  b.vel.vz += grav_bz
  return a, b
}

func ApplyVelocity(a Moon) Moon {
  a.pos.px += a.vel.vx
  a.pos.py += a.vel.vy
  a.pos.pz += a.vel.vz
  return a
}

func GetEnergy(a Moon) (int, int) {
  potential := int(math.Abs(float64(a.pos.px)) + math.Abs(float64(a.pos.py)) + math.Abs(float64(a.pos.pz)))
  kinetic := int(math.Abs(float64(a.vel.vx)) + math.Abs(float64(a.vel.vy)) + math.Abs(float64(a.vel.vz)))
  return potential, kinetic
}

func GetTotalEnergy(moon_list []Moon) int {
  total_energy := 0
  for _, moon := range moon_list {
    potential, kinetic := GetEnergy(moon)
    total_energy += potential * kinetic
  }
  return total_energy
}

func main() {
  file, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  moon_list := []Moon{}
  for scanner.Scan() {
    // By default, each token is separated by newline
    line := scanner.Text()
    line = strings.Trim(line, "<>")
  	line = strings.ReplaceAll(line, "x", "")
  	line = strings.ReplaceAll(line, "y", "")
  	line = strings.ReplaceAll(line, "z", "")
  	line = strings.ReplaceAll(line, "=", "")
  	line = strings.ReplaceAll(line, " ", "")
  	coords := strings.Split(line, ",")
  	x, _ := strconv.Atoi(coords[0])
    y, _ := strconv.Atoi(coords[1])
    z, _ := strconv.Atoi(coords[2])
    position := Position{x,y,z}
    velocity := Velocity{0,0,0}
    // Grab parameters from line and build initial positions for moon
    moon := Moon{position, velocity}
    moon_list = append(moon_list, moon)
  }
  //  Solution to Part 1
  // end_of_time := 20
  // for time := 0; time < end_of_time; time++ {
  //   // DO YOU UNDERSTAND THE GRAVITY OF THE SITUATION?
  //   for i := 0; i < len(moon_list) - 1; i++ {
  //     for j := i + 1; j < len(moon_list); j++ {
  //       moon_list[i], moon_list[j] = ApplyGravity(moon_list[i], moon_list[j])
  //     }
  //   }
  //   // GET A MOVE ON
  //   for i := 0; i < len(moon_list); i++ {
  //     moon_list[i] = ApplyVelocity(moon_list[i])
  //   }
  // }
  // fmt.Println(moon_list)
  // // fmt.Println(GetTotalEnergy(moon_list))

  // Part 2 - Without breaking out the calculus, we can do each dimension separately
  // Then record results, then do lcm of each plane
  initial_state := make([]Moon, 4)
  _ = copy(initial_state, moon_list)
  // Move time forward 1
  for i := 0; i < len(moon_list) - 1; i++ {
    for j := i + 1; j < len(moon_list); j++ {
      moon_list[i], moon_list[j] = ApplyGravity(moon_list[i], moon_list[j])
    }
  }
  for i := 0; i < len(moon_list); i++ {
    moon_list[i] = ApplyVelocity(moon_list[i])
  }
  fmt.Println(moon_list)
  fmt.Println(initial_state)
  equal := false
  for time := 1; !equal; time++ {
    // DO YOU UNDERSTAND THE GRAVITY OF THE SITUATION?
    for i := 0; i < len(moon_list) - 1; i++ {
      for j := i + 1; j < len(moon_list); j++ {
        moon_list[i], moon_list[j] = ApplyGravity(moon_list[i], moon_list[j])
      }
    }
    // GET A MOVE ON
    for i := 0; i < len(moon_list); i++ {
      moon_list[i] = ApplyVelocity(moon_list[i])
    }
    for i := 0; i < len(moon_list); i++ {
      if (moon_list[i].pos.pz == initial_state[i].pos.pz) && (moon_list[i].vel.vz == 0) {
        equal = true
      } else {
        equal = false
        break
      }
    }
    if equal {
      fmt.Println(time)
    }
  }
  fmt.Println(moon_list)

  // time_x := 161428
  // time_y := 231614
  // time_z := 116328
  // lcm_xy := ComputeLCM(time_x, time_y)
  // fmt.Println(lcm_xy)
  // lcm_xyz := ComputeLCM(lcm_xy, time_z)
  // fmt.Println(lcm_xyz)

}
