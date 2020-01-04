package main

import (
    "fmt"
    "math"
    // "io/ioutil"
    "strconv"
    "os"
    "bufio"
)

// Fuel = floor(mass / 3)- 2
func calculateFuel(mass float64) float64{
  return math.Floor(mass / 3) - 2
}

func main() {
  // Open file with input data.
  file, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  running_total := 0.0
  for scanner.Scan() {
    // By default, each token is separated by newline
    mass_text := scanner.Text()
    mass_int, _ := strconv.Atoi(mass_text)
    mass_float64 := float64(mass_int)
    fuel := calculateFuel(mass_float64)
    fuel_of_fuel := calculateFuel(fuel)
    for fuel_of_fuel > 0 {
      fuel = fuel + fuel_of_fuel
      fuel_of_fuel = calculateFuel(fuel_of_fuel)
    }
    running_total = running_total + fuel
    fmt.Println(fuel)
  }
  fmt.Println(int(running_total))
}
