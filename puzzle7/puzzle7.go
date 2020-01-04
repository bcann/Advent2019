package main

import (
    "fmt"
    "math"
)

const (
        POSITION = iota  //  0
        IMMEDIATE = iota  //  1
)

// one array for each amp in problem
var storage_I_guess = [6][]int{}
var memory_bank = []int{3,8,1001,8,10,8,105,1,0,0,21,46,55,72,85,110,191,272,353,434,99999,3,9,1002,9,5,9,1001,9,2,9,102,3,9,9,101,2,9,9,102,4,9,9,4,9,99,3,9,102,5,9,9,4,9,99,3,9,1002,9,2,9,101,2,9,9,1002,9,2,9,4,9,99,3,9,1002,9,4,9,101,3,9,9,4,9,99,3,9,1002,9,3,9,101,5,9,9,1002,9,3,9,101,3,9,9,1002,9,5,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,99}

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

func ConstructPhaseSequence(a int, b int, c int, d int, e int) int {
  phase_sequence := a * 10000
  phase_sequence += b * 1000
  phase_sequence += c * 100
  phase_sequence += d * 10
  phase_sequence += e
  return phase_sequence
}

func GetRemainders(unaltered []int, val int) []int {
  new_array := []int{}
  for _, x := range unaltered {
    if x != val {
	  new_array = append(new_array, x)
    }
  }
  // fmt.Println(new_array)
  return new_array
}

func RunProgram(input [2]int, inst_p int) (int, bool, int) {
  chain_input := input[0]
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
      StoreInstruction(instruction, inst_p, chain_input)
      chain_input = input[1]
    } else if opcode == 4 {
      output = LoadInstruction(instruction, inst_p)
      inst_p += inst_len
      return output, false, inst_p
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
      return output, true, inst_p
    }
    inst_p += inst_len
  }
  return output, true, inst_p // Shouldn't be reachable
}

// Copies contents of memory_bank into each storage array, fresh program in storage[0]]
func InitializeStorage() {
  for i := 0; i < len(storage_I_guess); i++ {
    storage_I_guess[i] = CopyIntSlice(memory_bank)
  }
}

// Used each time we try a new phase sequence
func ResetStorage() {
  for i := 1; i < len(storage_I_guess); i++ {
    storage_I_guess[i] = CopyIntSlice(storage_I_guess[0])
  }
}

func ResetMemory() {
  memory_bank = CopyIntSlice(storage_I_guess[0])
}

func LoadStorage(amp int) {
  memory_bank = CopyIntSlice(storage_I_guess[amp])
}

func StoreMemory(amp int) {
  storage_I_guess[amp] = CopyIntSlice(memory_bank)
}

func CopyIntSlice(slice []int) []int {
  copy := []int{}
  for _, x := range slice {
    copy = append(copy, x)
  }
  return copy
}

func main() {
  output := 0
  sequences := []int{5,6,7,8,9}
  max_output := 0
  // Enumerate all the possible phase sequences
  fmt.Println("PROGRAM TO DESTROY HUMANITY ---- INTIALIZE")
  InitializeStorage()
  for _, a := range sequences {
    remainders_a := GetRemainders(sequences, a)
    for _, b := range remainders_a {
      remainders_b := GetRemainders(remainders_a, b)
      for _, c := range remainders_b {
        remainders_c := GetRemainders(remainders_b, c)
        for _, d := range remainders_c {
          remainders_d := GetRemainders(remainders_c, d)
          for _, e := range remainders_d {
            //phase_sequence := ConstructPhaseSequence(a, b, c, d, e)
            amp_phase_map := map[int]int{
              5: a,
              6: b,
              7: c,
              8: d,
              9: e,
            }
            halt := false
            inst_p := [5]int{0,0,0,0,0}
            // Initial pass with phase sequence
            for i := 5; i <= 9; i++ {
              LoadStorage(i - 4)
              input := [2]int{amp_phase_map[i], output}
              output, halt, inst_p[i-5] = RunProgram(input, 0)
              StoreMemory(i-4)
            }
            for !halt {
              for i := 5; i <= 9; i++ {
                LoadStorage(i - 4)
                input := [2]int{output, output}
                output, halt, inst_p[i-5] = RunProgram(input, inst_p[i-5])
                if output > max_output { // This is a hack, but it gets the answer
                  max_output = output
                }
                StoreMemory(i-4)
              }
            }
			      output = 0
            ResetStorage() // Prepare for next phase sequence
            ResetMemory()
          }
        }
      }
    }
  }
  fmt.Println(max_output)
}
