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

// Not sure what capacity we need anymore
var memory_bank = [5000]int{1102,34463338,34463338,63,1007,63,34463338,63,1005,63,53,1102,1,3,1000,109,988,209,12,9,1000,209,6,209,3,203,0,1008,1000,1,63,1005,63,65,1008,1000,2,63,1005,63,904,1008,1000,0,63,1005,63,58,4,25,104,0,99,4,0,104,0,99,4,17,104,0,99,0,0,1102,1,21,1004,1101,28,0,1016,1101,0,27,1010,1102,36,1,1008,1102,33,1,1013,1101,0,22,1012,1101,0,37,1011,1102,34,1,1017,1102,466,1,1027,1102,1,484,1029,1102,1,699,1024,1102,1,1,1021,1101,0,0,1020,1102,1,24,1015,1101,0,473,1026,1101,653,0,1022,1102,26,1,1007,1102,25,1,1006,1101,0,39,1014,1102,646,1,1023,1101,690,0,1025,1102,1,29,1019,1101,32,0,1018,1101,30,0,1002,1101,0,20,1001,1102,1,38,1005,1102,1,23,1003,1101,0,31,1000,1101,35,0,1009,1101,0,493,1028,109,5,1208,0,37,63,1005,63,201,1001,64,1,64,1106,0,203,4,187,1002,64,2,64,109,-4,2107,36,8,63,1005,63,223,1001,64,1,64,1105,1,225,4,209,1002,64,2,64,109,18,21107,40,41,-9,1005,1010,243,4,231,1105,1,247,1001,64,1,64,1002,64,2,64,109,6,21107,41,40,-9,1005,1016,267,1001,64,1,64,1106,0,269,4,253,1002,64,2,64,109,-19,21102,42,1,5,1008,1011,42,63,1005,63,291,4,275,1105,1,295,1001,64,1,64,1002,64,2,64,109,15,1205,0,309,4,301,1105,1,313,1001,64,1,64,1002,64,2,64,109,-27,2101,0,9,63,1008,63,20,63,1005,63,333,1106,0,339,4,319,1001,64,1,64,1002,64,2,64,109,19,21102,43,1,6,1008,1019,45,63,1005,63,363,1001,64,1,64,1105,1,365,4,345,1002,64,2,64,109,1,21108,44,47,-3,1005,1011,385,1001,64,1,64,1106,0,387,4,371,1002,64,2,64,109,-22,1201,9,0,63,1008,63,21,63,1005,63,411,1001,64,1,64,1106,0,413,4,393,1002,64,2,64,109,9,1207,0,19,63,1005,63,433,1001,64,1,64,1106,0,435,4,419,1002,64,2,64,109,-9,2107,30,8,63,1005,63,453,4,441,1105,1,457,1001,64,1,64,1002,64,2,64,109,25,2106,0,10,1001,64,1,64,1106,0,475,4,463,1002,64,2,64,109,11,2106,0,0,4,481,1001,64,1,64,1105,1,493,1002,64,2,64,109,-18,2108,21,-6,63,1005,63,511,4,499,1106,0,515,1001,64,1,64,1002,64,2,64,109,-12,2108,18,6,63,1005,63,535,1001,64,1,64,1106,0,537,4,521,1002,64,2,64,109,19,21101,45,0,-7,1008,1010,45,63,1005,63,563,4,543,1001,64,1,64,1105,1,563,1002,64,2,64,109,-10,1207,-5,31,63,1005,63,581,4,569,1106,0,585,1001,64,1,64,1002,64,2,64,109,-8,2102,1,5,63,1008,63,21,63,1005,63,611,4,591,1001,64,1,64,1105,1,611,1002,64,2,64,109,5,1201,0,0,63,1008,63,21,63,1005,63,633,4,617,1106,0,637,1001,64,1,64,1002,64,2,64,109,13,2105,1,6,1001,64,1,64,1106,0,655,4,643,1002,64,2,64,109,-7,1202,-3,1,63,1008,63,26,63,1005,63,681,4,661,1001,64,1,64,1106,0,681,1002,64,2,64,109,12,2105,1,2,4,687,1001,64,1,64,1105,1,699,1002,64,2,64,109,-28,1208,8,30,63,1005,63,717,4,705,1106,0,721,1001,64,1,64,1002,64,2,64,109,10,1202,1,1,63,1008,63,40,63,1005,63,745,1001,64,1,64,1105,1,747,4,727,1002,64,2,64,109,10,21108,46,46,-2,1005,1012,765,4,753,1105,1,769,1001,64,1,64,1002,64,2,64,109,-2,1205,8,781,1106,0,787,4,775,1001,64,1,64,1002,64,2,64,109,-9,2101,0,0,63,1008,63,23,63,1005,63,809,4,793,1105,1,813,1001,64,1,64,1002,64,2,64,109,9,1206,8,831,4,819,1001,64,1,64,1106,0,831,1002,64,2,64,109,-9,2102,1,-2,63,1008,63,22,63,1005,63,855,1001,64,1,64,1106,0,857,4,837,1002,64,2,64,109,4,21101,47,0,10,1008,1017,50,63,1005,63,877,1105,1,883,4,863,1001,64,1,64,1002,64,2,64,109,18,1206,-4,895,1105,1,901,4,889,1001,64,1,64,4,64,99,21101,0,27,1,21102,915,1,0,1106,0,922,21201,1,56639,1,204,1,99,109,3,1207,-2,3,63,1005,63,964,21201,-2,-1,1,21102,1,942,0,1106,0,922,22102,1,1,-1,21201,-2,-3,1,21101,0,957,0,1106,0,922,22201,1,-1,-2,1106,0,968,22102,1,-2,-2,109,-3,2106,0,0}

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
  fmt.Println(store_index)
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

func RunProgram(input int, inst_p int) int {
  output := 0
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
      output = LoadInstruction(instruction, inst_p)
      fmt.Println(output)
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
      return output
    }
    if jumped {
      continue
    }
    inst_p += inst_len
  }
  return output  // Shouldn't be reachable
}

func main() {
  output := 0
  input := 2
  inst_p := 0
  output = RunProgram(input, inst_p)
  fmt.Println(output)
}
