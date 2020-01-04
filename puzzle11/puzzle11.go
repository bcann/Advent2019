package main

import (
    "fmt"
    "math"
)

const (
        POSITION = iota  //  0
        IMMEDIATE = iota  //  1
        RELATIVE = iota // 2
)

const (
       NORTH = iota  //  0
       EAST = iota  //  1
       SOUTH = iota // 2
       WEST = iota // 3
)

// Not sure what capacity we need anymore
var memory_bank = [5000]int{3,8,1005,8,332,1106,0,11,0,0,0,104,1,104,0,3,8,102,-1,8,10,101,1,10,10,4,10,108,1,8,10,4,10,101,0,8,28,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,51,1,1103,5,10,1,1104,9,10,2,1003,0,10,1,5,16,10,3,8,102,-1,8,10,101,1,10,10,4,10,108,0,8,10,4,10,1001,8,0,88,1006,0,2,1006,0,62,2,8,2,10,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,1,10,4,10,102,1,8,121,1006,0,91,1006,0,22,1006,0,23,1006,0,1,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,155,1006,0,97,1,1004,2,10,2,1003,6,10,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,1002,8,1,187,1,104,15,10,2,107,9,10,1006,0,37,1006,0,39,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,0,8,10,4,10,102,1,8,223,2,2,17,10,1,1102,5,10,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,1001,8,0,253,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1002,8,1,276,1006,0,84,3,8,102,-1,8,10,101,1,10,10,4,10,1008,8,0,10,4,10,1001,8,0,301,2,1009,9,10,1006,0,10,2,102,15,10,101,1,9,9,1007,9,997,10,1005,10,15,99,109,654,104,0,104,1,21102,1,936995738516,1,21101,0,349,0,1105,1,453,21102,1,825595015976,1,21102,1,360,0,1105,1,453,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21102,46375541763,1,1,21101,0,407,0,1105,1,453,21102,1,179339005019,1,21101,0,418,0,1106,0,453,3,10,104,0,104,0,3,10,104,0,104,0,21102,825012036372,1,1,21102,441,1,0,1105,1,453,21101,988648461076,0,1,21101,452,0,0,1105,1,453,99,109,2,22102,1,-1,1,21102,40,1,2,21102,484,1,3,21101,0,474,0,1106,0,517,109,-2,2105,1,0,0,1,0,0,1,109,2,3,10,204,-1,1001,479,480,495,4,0,1001,479,1,479,108,4,479,10,1006,10,511,1102,1,0,479,109,-2,2105,1,0,0,109,4,2102,1,-1,516,1207,-3,0,10,1006,10,534,21101,0,0,-3,21202,-3,1,1,22101,0,-2,2,21102,1,1,3,21102,553,1,0,1106,0,558,109,-4,2106,0,0,109,5,1207,-3,1,10,1006,10,581,2207,-4,-2,10,1006,10,581,22102,1,-4,-4,1105,1,649,21202,-4,1,1,21201,-3,-1,2,21202,-2,2,3,21101,0,600,0,1105,1,558,21201,1,0,-4,21101,0,1,-1,2207,-4,-2,10,1006,10,619,21101,0,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,641,22102,1,-1,1,21102,1,641,0,106,0,516,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2105,1,0}

var relative_base = 0

var inst_len_mapping = map[int]int{
  1:  4,
  2:  4,
  3:  2,
  4:  2,
  5:  3,
  6:  3,
  7:  4,
  8:  4,
  9:  2,
  99: 0,
}

// Function gets the mode of a parameter
func GetMode(instruction int, param int) int {
  // Move parameter digit to the ones digit, then mod 10 to isolate it
   mode := (instruction / int(math.Pow10(param+1))) % 10
   return mode
}

func GetOpcode(instruction int) int {
   // isolate last 2 digits
   opcode := instruction % 100
   return opcode
}

func GetValWithMode(inst_p int, mode int, param int) int {
  if mode == IMMEDIATE {
    return memory_bank[inst_p + param]
  } else if mode == RELATIVE {
    return memory_bank[memory_bank[inst_p + param] + relative_base]
  } else {
    return memory_bank[memory_bank[inst_p + param]]
  }
}

func GetAddrWithMode(inst_p int, mode int, param int) int {
  if mode == IMMEDIATE { // This shouldn't happen
    fmt.Println("This shouldn't happen")
    return memory_bank[inst_p + param]
  } else if mode == RELATIVE {
    return memory_bank[inst_p + param] + relative_base
  } else {
    return memory_bank[inst_p + param]
  }
}

func AddInstruction(instruction int, inst_p int) bool {
  // num_params := len(strconv.Itoa(instruction)) - 2
  // for now, only using 2 params
  max_params := 2
  add_total := 0
  for param := 1; param <= max_params; param++ {
    add_total += GetValWithMode(inst_p, GetMode(instruction, param), param)
  }
  store_index := GetAddrWithMode(inst_p, GetMode(instruction, 3), 3)
  memory_bank[store_index] = add_total
  return true
}

func MultiplyInstruction(instruction int, inst_p int) bool {
  // num_params := len(strconv.Itoa(instruction)) - 2
  // for now, only using 2 params
  max_params := 2
  mult_total := 1
  for param := 1; param <= max_params; param++ {
    mult_total *= GetValWithMode(inst_p, GetMode(instruction, param), param)
  }
  store_index := GetAddrWithMode(inst_p, GetMode(instruction, 3), 3)
  memory_bank[store_index] = mult_total
  return true
}

// 1 parameter simplifies this
func StoreInstruction(instruction int, inst_p int, input int) bool {
  store_index := GetAddrWithMode(inst_p, GetMode(instruction, 1), 1)
  memory_bank[store_index] = input
  return true
}

// 1 parameter simplifies this
func LoadInstruction(instruction int, inst_p int) int {
  output := GetValWithMode(inst_p, GetMode(instruction, 1), 1)
  return output
}

func JumpTrueInstruction(instruction int, inst_p int) (int, bool) {
  compare_val := GetValWithMode(inst_p, GetMode(instruction, 1), 1)
  jumped := compare_val != 0
  if jumped {
    inst_p = GetValWithMode(inst_p, GetMode(instruction, 2), 2)
  }
  return inst_p, jumped
}

func JumpFalseInstruction(instruction int, inst_p int) (int, bool) {
  compare_val := GetValWithMode(inst_p, GetMode(instruction, 1), 1)
  jumped := compare_val == 0
  if jumped {
    inst_p = GetValWithMode(inst_p, GetMode(instruction, 2), 2)
  }
  return inst_p, jumped
}

func LessThanInstruction(instruction int, inst_p int) bool {
  max_params := 2
  compare_params := []int{}
  for param := 1; param <= max_params; param++ {
      compare_params = append(compare_params, GetValWithMode(inst_p, GetMode(instruction, param), param))
  }
  result := compare_params[0] < compare_params[1]
  store_index := GetAddrWithMode(inst_p, GetMode(instruction, 3), 3)
  if result == true {
    memory_bank[store_index] = 1
  } else {
    memory_bank[store_index] = 0
  }
  return true
}

func EqualsInstruction(instruction int, inst_p int) bool {
  max_params := 2
  compare_params := []int{}
  for param := 1; param <= max_params; param++ {
      compare_params = append(compare_params, GetValWithMode(inst_p, GetMode(instruction, param), param))
  }
  result := compare_params[0] == compare_params[1]
  store_index := GetAddrWithMode(inst_p, GetMode(instruction, 3), 3)
  if result == true {
    memory_bank[store_index] = 1
  } else {
    memory_bank[store_index] = 0
  }
  return true
}

func RelativeInstruction(instruction int, inst_p int) bool {
  val := GetValWithMode(inst_p, GetMode(instruction, 1), 1)
  relative_base += val
  return true
}

func RunProgram(input int, inst_p int) ([]int, bool, int) {
  output := []int{}
  for inst_p < len(memory_bank) {
    instruction := memory_bank[inst_p]
    opcode := GetOpcode(instruction)
    inst_len := inst_len_mapping[opcode]
    jumped := false
    if opcode == 1 {
      AddInstruction(instruction, inst_p)
    } else if opcode == 2 {
      MultiplyInstruction(instruction, inst_p)
    } else if opcode == 3 {
      StoreInstruction(instruction, inst_p, input)
    } else if opcode == 4 {
      temp_output := LoadInstruction(instruction, inst_p)
      output = append(output, temp_output)
    } else if opcode == 5 {
      inst_p, jumped = JumpTrueInstruction(instruction, inst_p)
    } else if opcode == 6 {
      inst_p, jumped = JumpFalseInstruction(instruction, inst_p)
    } else if opcode == 7 {
      LessThanInstruction(instruction, inst_p)
    } else if opcode == 8 {
      EqualsInstruction(instruction, inst_p)
    } else if opcode == 9 {
      RelativeInstruction(instruction, inst_p)
    } else if opcode == 99 {
      return output, true, inst_p
    }
    if jumped {
      continue
    }
    inst_p += inst_len
    if len(output) == 2 {
      return output, false, inst_p
    }
  }
  return output, true, inst_p  // Shouldn't be reachable
}

type Tile struct {
  X int
  Y int
}

func ChangeOrientation(rotation int, orientation int) int {
  new_orientation := orientation
  if rotation == 0 {
    new_orientation -= 1
    if new_orientation < 0 {
      new_orientation += 4
    }
  } else {
    new_orientation += 1
    if new_orientation > 3 {
      new_orientation -= 4
    }
  }
  return new_orientation
}

func MoveTile(orientation int, current_tile Tile) Tile {
  if orientation == NORTH {
    return Tile{current_tile.X, current_tile.Y+1}
  } else if orientation == EAST {
    return Tile{current_tile.X+1, current_tile.Y}
  } else if orientation == SOUTH {
    return Tile{current_tile.X, current_tile.Y-1}
  } else { // weast
    return Tile{current_tile.X-1, current_tile.Y}
  }
}

// Do a dry run, collecting all outputs first
// Then build robot controller.
func main() {
  output := []int{}
  input := 1
  inst_p := 0
  halt := false
  current_tile := Tile{0,0}
  orientation := NORTH
  visited_tiles := make(map[Tile]bool)
  for  !halt {
    output, halt, inst_p = RunProgram(input, inst_p)
    if !halt {
      orientation = ChangeOrientation(output[1], orientation)
      visited_tiles[current_tile] = output[0] == 1
      current_tile = MoveTile(orientation, current_tile)
      if color, ok := visited_tiles[current_tile]; ok { // ISNT GOLANG COOL???
        if color { // white
          input = 1
        } else { // black
          input = 0
        }
      } else { // First time?
        input = 0
      }
    }
  }
  print(len(visited_tiles)) // Part 1 solution
  // Find leftmost, rightmost, northmost, southmost coordinates
  out := 0
  over := 0
  up := 0
  down := 0
  for k, _ := range visited_tiles {
    fmt.Println(k)
    if k.X > over {
      over = k.X
    } else if k.X < out {
      out = k.X
    }
    if k.Y > up {
      up = k.Y
    } else if k.Y < down {
      down = k.Y
    }
  }

  fmt.Printf("Over: %d Out: %d Up: %d Down: %d\n", over, out, up, down)
  for y := up; y >= down; y-- {
    fmt.Printf("\n")
    for x := out; x <= over; x++ {
      if color, ok := visited_tiles[Tile{x,y}]; ok { // ISNT GOLANG COOL???
        if color { // white
          fmt.Printf("#")
        } else { // black
          fmt.Printf(".")
        }
      } else {
        fmt.Printf(".")
      }
    }
  }
  // Sort all tiles by X, then sort by Y
}
