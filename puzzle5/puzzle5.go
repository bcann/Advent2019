package main

import (
    "fmt"
    "math"
)

const (
        POSITION = iota  //  0
        IMMEDIATE = iota  //  1
)

var memory_bank = []int{3,225,1,225,6,6,1100,1,238,225,104,0,1102,91,92,225,1102,85,13,225,1,47,17,224,101,-176,224,224,4,224,1002,223,8,223,1001,224,7,224,1,223,224,223,1102,79,43,225,1102,91,79,225,1101,94,61,225,1002,99,42,224,1001,224,-1890,224,4,224,1002,223,8,223,1001,224,6,224,1,224,223,223,102,77,52,224,1001,224,-4697,224,4,224,102,8,223,223,1001,224,7,224,1,224,223,223,1101,45,47,225,1001,43,93,224,1001,224,-172,224,4,224,102,8,223,223,1001,224,1,224,1,224,223,223,1102,53,88,225,1101,64,75,225,2,14,129,224,101,-5888,224,224,4,224,102,8,223,223,101,6,224,224,1,223,224,223,101,60,126,224,101,-148,224,224,4,224,1002,223,8,223,1001,224,2,224,1,224,223,223,1102,82,56,224,1001,224,-4592,224,4,224,1002,223,8,223,101,4,224,224,1,224,223,223,1101,22,82,224,1001,224,-104,224,4,224,1002,223,8,223,101,4,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,8,226,677,224,102,2,223,223,1005,224,329,1001,223,1,223,1007,226,226,224,1002,223,2,223,1006,224,344,101,1,223,223,108,226,226,224,1002,223,2,223,1006,224,359,1001,223,1,223,107,226,677,224,102,2,223,223,1006,224,374,101,1,223,223,8,677,677,224,102,2,223,223,1006,224,389,1001,223,1,223,1008,226,677,224,1002,223,2,223,1006,224,404,101,1,223,223,7,677,677,224,1002,223,2,223,1005,224,419,101,1,223,223,1108,226,677,224,1002,223,2,223,1005,224,434,101,1,223,223,1108,226,226,224,102,2,223,223,1005,224,449,1001,223,1,223,107,226,226,224,102,2,223,223,1005,224,464,101,1,223,223,1007,677,677,224,102,2,223,223,1006,224,479,101,1,223,223,1007,226,677,224,102,2,223,223,1005,224,494,1001,223,1,223,1008,226,226,224,1002,223,2,223,1005,224,509,1001,223,1,223,1108,677,226,224,1002,223,2,223,1006,224,524,1001,223,1,223,108,677,677,224,1002,223,2,223,1005,224,539,101,1,223,223,108,226,677,224,1002,223,2,223,1005,224,554,101,1,223,223,1008,677,677,224,1002,223,2,223,1006,224,569,1001,223,1,223,1107,677,677,224,102,2,223,223,1005,224,584,1001,223,1,223,7,677,226,224,102,2,223,223,1005,224,599,1001,223,1,223,8,677,226,224,1002,223,2,223,1005,224,614,1001,223,1,223,7,226,677,224,1002,223,2,223,1006,224,629,101,1,223,223,1107,677,226,224,1002,223,2,223,1005,224,644,1001,223,1,223,1107,226,677,224,102,2,223,223,1006,224,659,1001,223,1,223,107,677,677,224,1002,223,2,223,1005,224,674,101,1,223,223,4,223,99,226}
// var memory_bank = []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
// 1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
// 999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}

var inst_len_mapping = map[int]int{
  1:  4,
  2:  4,
  3:  2,
  4:  2,
  5:  3,
  6:  3,
  7:  4,
  8:  4,
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
  } else {
    return memory_bank[memory_bank[inst_p + param]]
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
  memory_bank[memory_bank[inst_p+max_params+1]] = add_total
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
  memory_bank[memory_bank[inst_p+max_params+1]] = mult_total
  return true
}

// 1 parameter simplifies this
func StoreInstruction(instruction int, inst_p int, input int) bool {
  store_index := memory_bank[inst_p+1] // literally just first parameter
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
  if result == true {
    memory_bank[memory_bank[inst_p+3]] = 1
  } else {
    memory_bank[memory_bank[inst_p+3]] = 0
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
  if result == true {
    memory_bank[memory_bank[inst_p+3]] = 1
  } else {
    memory_bank[memory_bank[inst_p+3]] = 0
  }
  return true
}

func main() {
  // Read input into one array. Chunk if we need bigger.
  inst_p := 0
  input := 5
  output := 0
  for inst_p < len(memory_bank) {
    instruction := memory_bank[inst_p]
    opcode := GetOpcode(instruction)
      // fmt.Println(inst_p)
      // fmt.Println(instruction)
    inst_len := inst_len_mapping[opcode]
    jumped := false
    if opcode == 1 {
      AddInstruction(instruction, inst_p)
    } else if opcode == 2 {
      MultiplyInstruction(instruction, inst_p)
    } else if opcode == 3 {
      StoreInstruction(instruction, inst_p, input)
    } else if opcode == 4 {
      output = LoadInstruction(instruction, inst_p)
    } else if opcode == 5 {
      inst_p, jumped = JumpTrueInstruction(instruction, inst_p)
      if jumped {
        continue
      }
    } else if opcode == 6 {
      inst_p, jumped = JumpFalseInstruction(instruction, inst_p)
      if jumped {
        continue
      }
    } else if opcode == 7 {
      LessThanInstruction(instruction, inst_p)
    } else if opcode == 8 {
      EqualsInstruction(instruction, inst_p)
    } else if opcode == 99 {
      break
    }
    inst_p += inst_len
  }
  fmt.Println(output)
}
